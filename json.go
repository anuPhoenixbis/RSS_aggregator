package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
func respondWithJSON(w http.ResponseWriter , code int , payload interface{}){//payload interface holds the data
	// converting json to bytes so that we write it directly into an http reponse
	data , err := json.Marshal(payload)//It takes an input of type interface{} (which means it can accept any type) and converts it into a JSON-encoded byte slice.
// checking for error while converting
	if err!=nil{
		log.Println("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)//WriteHeader sends an HTTP response header with the provided status code.
		return
	}
	//if not error faced
//Adding header metadata
	//Add adds the key, value pair to the header. It appends to any existing values associated with key.
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(200)//successful response
	//writing byte again to json
	w.Write(data)//Write writes the data to the connection as part of an HTTP response.
}