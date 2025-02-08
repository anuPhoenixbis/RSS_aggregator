package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	err:= godotenv.Load(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }
// 	The os.Getenv function retrieves the value of the environment variable named "PORT".
// If the environment variable is not set, os.Getenv returns an empty string.
	portString := os.Getenv("PORT")
	//checking of the port
	if portString == "" {
		log.Fatal("Port is not found in the enviornment")
	}

	// An HTTP router is a component within a web application that determines which specific piece of code should
	//  handle an incoming HTTP request based on the URL path
	//router is our new router obj
	//import these : "github.com/go-chi/chi" and "github.com/go-chi/cors"  
	router := chi.NewRouter()
	

	//to accept request from the browser, reponse is not handled here
	// This applies a middleware to the router that handles CORS rules.
	// cors.Handler() is used to set up CORS policies.
	// cors.Options{} defines the specific CORS rules.
	router.Use(cors.Handler(cors.Options{
// 		This allows requests from any origin that starts with https:// or http://.
// The * acts as a wildcard, meaning any domain is accepted.
// This is useful when serving an API that needs to be accessed from multiple front-end applications.
		AllowedOrigins:   []string{"https://*", "http://*"},
		// Specifies which HTTP methods are allowed when making cross-origin requests.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Allows all headers in incoming requests.
		AllowedHeaders:   []string{"*"},
		// Specifies which headers the browser can access in the response.
		ExposedHeaders:   []string{"Link"},
		// This means the server does not allow credentials (such as cookies, authentication tokens, or session identifiers) in cross-origin requests.
		AllowCredentials: false,
		// Defines the maximum time (in seconds) that the results of a preflight request (OPTIONS request) can be cached by the browser.
		MaxAge:           300,
	}))
	
	//creating a server
// 	This line declares a variable srv and initializes it with a pointer to a new http.Server struct.
// The & operator is used to get the address of the http.Server struct, creating a pointer to it.
	srv := &http.Server{
		// The Handler is responsible for handling incoming HTTP requests. In this case, the router will handle the routing 
		// of requests to the appropriate handlers.
		Handler: router,
		// The Handler is responsible for handling incoming HTTP requests. In this case, the router will handle the routing 
		// of requests to the appropriate handlers.
		Addr:    ":" + portString,//address of the server/site
	}
	log.Printf("Server is starting on port %v" , portString)
	err = srv.ListenAndServe()//if any kind of error occurs server gets activated it gets stored in this var
	//if error occurs then log out
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Port: ",portString)
}