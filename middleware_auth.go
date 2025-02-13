package main

import (
	"fmt"
	"net/http"

	"github.com/anuPhoenixbis/RSS_Agg/internal/auth"
	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
)

//defining the type of the authHandler similar to the http.HandlerFunc but with an extra parameter of database.User
type authHandler func(http.ResponseWriter , *http.Request , database.User)

//to make it passable to the http.HandlerFunc we need to remove the extra parameter
func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(*r) // getting api key here

		// checking for errors
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}