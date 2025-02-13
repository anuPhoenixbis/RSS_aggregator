package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request , user database.User) {
    type parameter struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400 , fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	//this create user method was created by sqlc in go db by reading from the sql  
	feed , err := apiConfig.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err!=nil{
		respondWithError(w, 400 ,fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w,201,databaseFeedToFeed(feed))//changing the user in the db with our user from models.go
}


func (apiConfig *apiConfig)handlerGetFeed(w http.ResponseWriter, r *http.Request ) {
    
	//this create user method was created by sqlc in go db by reading from the sql  
	feed , err := apiConfig.DB.GetFeeds(r.Context())
	if err!=nil{
		respondWithError(w, 400 ,fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w,201,databaseFeedsToFeeds(feed))//changing the user in the db with our user from models.go
}
//respondWithJSON is a helper function to respond with a JSON payload