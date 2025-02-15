package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"
)

func (apiConfig *apiConfig)handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request , user database.User) {
    type parameter struct{
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400 , fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	//this create user method was created by sqlc in go db by reading from the sql  
	feedfollow , err := apiConfig.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err!=nil{
		respondWithError(w, 400 ,fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w,201,databaseFeedFollowToFeedFollow(feedfollow))//changing the user in the db with our user from models.go
}


func (apiConfig *apiConfig)handlerGetFeedFollow(w http.ResponseWriter, r *http.Request , user database.User)  {
	//this create user method was created by sqlc in go db by reading from the sql  
	feedfollows , err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	
	if err!=nil{
		respondWithError(w, 400 ,fmt.Sprintf("Couldn't get feed follow: %v", err))
		return
	}

	respondWithJSON(w,201,databaseFeedFollowsToFeedFollows(feedfollows))//changing the user in the db with our user from models.go
}


func (apiConfig *apiConfig)handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request , user database.User)  {
	feedFollowIdStr := chi.URLParam(r,"feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIdStr)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}
	err = apiConfig.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w , 200 , struct{}{})//feed delete successfully
}

