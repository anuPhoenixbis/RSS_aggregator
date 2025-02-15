package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/google/uuid"
)

//The purpose of this code is to periodically scrape RSS feeds from a database using a specified number of goroutines
func StartScraping(
	db *database.Queries,//the connection to the database
	concurrency int,//the number of goroutines to use for scraping
	timeBetweenRequest time.Duration,//the time to wait between requests
){
	log.Printf("Scraping on %v goroutines every %s duration",concurrency,timeBetweenRequest)
	// This is a receive operation on the channel ticker.C. The time.Ticker object has a channel C that sends the current time at intervals specified by the ticker. The loop waits for a value to be received from ticker.C before each iteration.
	ticker := time.NewTicker(timeBetweenRequest)//It initializes a time.Ticker that will trigger at intervals specified by timeBetweenRequest.
	for ; ; <-ticker.C{//The function enters an infinite loop that triggers on each tick of the ticker.
		feeds, err := db.GetNextFeedsToFetch(//It calls db.GetNextFeedsToFetch to retrieve the next set of feeds to fetch, passing a background context and the concurrency level.
			context.Background() ,
			 int32(concurrency) ,
			)
	if err !=nil { 
		log.Printf("Error getting feeds to fetch: %v",err)
		continue//should continue to the next iteration of the loop rather than exiting the loop
	}

	//this waitgroup is basically used to spawn say , 30 different goroutines and scrape them till all of them are done and then move to the next set of feeds
	wg :=  sync.WaitGroup{}//It creates a new sync.WaitGroup object to synchronize the goroutines.
	for _,feed := range feeds{
		wg.Add(1)//It increments the WaitGroup counter by 1.

		go scrapeFeed(db ,&wg , feed)//It starts a new goroutine to scrape the feed.
	}
	//until all the feeds are scraped we don't want to move to the next set of feeds
	wg.Wait()//It blocks until the WaitGroup counter is zero, which means all the goroutines have completed.
	}
}

func scrapeFeed(db *database.Queries , wg *sync.WaitGroup , feed database.Feed){
	defer wg.Done()//It decrements the WaitGroup counter by 1 when the goroutine completes.
	//scraping logic here

	//It calls db.MarkFeedAsFetched to mark the feed as fetched in the database.
	_,err := db.MarkFeedAsFetched(context.Background(),feed.ID)//It calls db.MarkFeedAsFetched to mark the feed as fetched in the database.
	if err != nil{
		log.Printf("Error marking feed as fetched: %v",err)
		return
	}

	//It fetches the feed from the URL 
	rssFeed , err := urlToFeed(feed.Url)//It calls the urlToFeed function to fetch the feed from the URL.
	if err != nil{
		log.Printf("Error fetching feed: %v",err)
		return
	}
	//It logs the number of posts found in the feed.
	for _,item := range rssFeed.Channel.Items{
		//It creates a new description object with a valid string and sets the description to the item.Description if it is not empty.
		description := sql.NullString{}
		if(item.Description != ""){//if there is no description then it will be null
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)//It parses the item.PubDate string into a time.Time object.
		if err != nil{
			log.Printf("couldn't parse date %v with error %v",item.PubDate,err)
			continue
		}

		_,err = db.CreatePost(context.Background(),database.CreatePostParams{
			ID : uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title : item.Title,
			Description : description,//passing the above created description
			PublishedAt: pubAt,
			Url : item.Link,
			FeedID : feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(),"duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}
	}
	log.Printf("Feed %s collected , %v posts found",feed.Name,len(rssFeed.Channel.Items))
}

