package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anuPhoenixbis/RSS_Agg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
// The type of the DB field is *database.Queries, which is a pointer to a Queries struct defined in the database package.
	DB *database.Queries//holds the connection to the db
}

func main(){

	err:= godotenv.Load(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }
// 	The os.Getenv function retrieves the value of the environment variable named "PORT".
// If the environment variable is not set, os.Getenv returns an empty string.

//server connection
	portString := os.Getenv("PORT")
	//checking of the port
	if portString == "" {
		log.Fatal("Port is not found in the enviornment")
	}

	//db code goes here
//db connection
	dbURL := os.Getenv("DB_URL")
	//checking of the port
	if portString == "" {
		log.Fatal("DB URL is not found in the enviornment")
	}

	//connecting the db and error handling
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	//getting queries from the connection  
	queries := database.New(conn)
	
	apiCfg := apiConfig{
		DB:queries,
	}
	
	defer conn.Close()
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


	
	//to hook up our handler_readiness to the main router
	v1Router := chi.NewRouter()
	
	//HandleFunc() gives output on all the request types(get,post,delete,etc) so we set it to Get only
	// v1Router.HandleFunc("/ready", handlerReadiness) sets up a route that listens for HTTP requests at the path /ready.
	// When a request is made to /ready, the handlerReadiness function will be called to handle the request.
	v1Router.Get("/ready", handlerReadiness)
	
	
	//error handler
	v1Router.Get("/err" , handlerError)
	
	//handles the user creation from the db
	v1Router.Post("/users" , apiCfg.handlerCreateUser) 

	//handles the user getting from the db
	// passing the getuserhandler in the middlewareAuth
	v1Router.Get("/users" , apiCfg.middlewareAuth(apiCfg.handlerGetUser))//passing the middleware here for the error handling

	//for the feed creation
	v1Router.Post("/feeds" , apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))//passing the middleware here for the error handling

	//for the feed getting
	v1Router.Get("/feeds" , apiCfg.handlerGetFeed)//passing the middleware here for the error handling


	//for feed follow
	v1Router.Post("/feed_follows" , apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))//passing the middleware here for the error handling


	//for feed follow getting
	v1Router.Get("/feed_follows" , apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))//passing the middleware here for the error handling

	//for feed follow delete
	v1Router.Delete("/feed_follows/{feedFollowID}" , apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))//passing the middleware here for the error handling
	
	//adding the main v1Router to the main router 
	// 	This line mounts the v1Router onto the main router at the path /v1.
	// This means that any routes defined in v1Router will be accessible under the /v1 path.
	router.Mount("/v1",v1Router)


	
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