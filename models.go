package main

import (
	"time"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/google/uuid"
)
// Each field has a corresponding JSON tag (e.g., json:"id") that specifies the key to use when the struct is serialized to JSON.
type User struct{
	ID 			uuid.UUID `json"id"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
	Name		string	  `json:"name"`
	APIKey      string	  `json:"api_key"`
}
type Feed struct{
	ID        uuid.UUID `json"id"`
	CreatedAt time.Time `json"created_at"`
	UpdatedAt time.Time `json"updated_at"`
	Name      string `json"name"`
	Url       string `json"url"`
	UserID    uuid.UUID `json"user_id"`
}

// This function converts a database.User struct to a User struct.
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID : dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name : dbUser.Name,
		APIKey : dbUser.ApiKey,
	}
}
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID : dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name : dbFeed.Name,
		Url :  dbFeed.Url,
		UserID : dbFeed.UserID,
	}
}
//takes a slice of database.User and returns a slice of User
func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}//creating an empty slice of Feed
	for _,dbFeed := range dbFeed{//iterating over the slice of database.User
		feeds = append(feeds,databaseFeedToFeed(dbFeed))//appending the converted User to the slice of User
	}
	return feeds
}