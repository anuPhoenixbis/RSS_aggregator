package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

//structure of rss feed xml file
type RSSFeed struct{
	Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		Items []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		PubDate string `xml:"pubDate"`
}


func urlToFeed(url string) (RSSFeed , error) {
	//creating a client to make a request to the url
	httpClient := http.Client{
		Timeout: 10*time.Second,//leaves the feed if it takes more than 10 seconds
	}

	resp,err := httpClient.Get(url)//getting the response here
	if err != nil{//error for not getting the response
		return RSSFeed{}, err
	}
	defer resp.Body.Close()//closing the body of the response
	
	data , err := io.ReadAll(resp.Body)//fetching the data from the response
	if err != nil{//error for not getting the data
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}//creating an empty rss feed to store the data
	err = xml.Unmarshal(data,&rssFeed)//unmarshalling/converting the data to the rss feed from xml , similar to json.Unmarshal
	if err != nil{//error for not unmarshalling the data
		return RSSFeed{}, err
	}
	return rssFeed, nil//returning the rss feed data converted from the xml
}