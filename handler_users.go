package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type parameter struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400 , fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	//this create user method was created by sqlc in go db by reading from the sql  
	user ,err := apiConfig.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,

	})
	if err!=nil{
		respondWithError(w, 400 ,fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w,201,databaseUserToUser(user))//changing the user in the db with our user from models.go
}

func (apiConfig *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request , user database.User) {
	respondWithJSON(w , 200 , databaseUserToUser(user))
}