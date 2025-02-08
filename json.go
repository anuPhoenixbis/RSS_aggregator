package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
//code int: This is the HTTP status code to be sent in the response.
//msg string: This is the error message to be included in the response.

//generate response with error
func respondWithError(w http.ResponseWriter , code int  , msg string){
	//This checks if the HTTP status code is 500 or greater (indicating a server error).
	if code>499{
		log.Println("Responding with 5XX error : " ,msg)
	}
	//The struct tag json:"error" specifies that the Error field should be serialized to JSON with the key "error".
	type errResponse struct{
		Error string `json:"error"`
	}

	//we will send a struct type of errResponse to respondWtihJSON method which should return the 
	//json given below
	// {
	// 	"error" : "something went wrong"
	// }

	//if not error is found at last call respond with json method
	respondWithJSON(w, code , errResponse{
		Error : msg ,
 	})
	//The respondWithJSON function is responsible for converting the errResponse struct to JSON and writing it to the HTTP response.
}

//generate response with json
func respondWithJSON(w http.ResponseWriter , code int , payload interface{}){//payload interface holds the data
	// converting json to bytes so that we write it directly into an http reponse
	data , err := json.Marshal(payload)//It takes an input of type interface{} (which means it can accept any type) and converts it into a JSON-encoded byte slice.
// checking for error while converting
	if err!=nil{
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)//WriteHeader sends an HTTP response header with the provided status code.
		return
	}
	//if not error faced
//Adding header metadata
	//Add adds the key, value pair to the header. It appends to any existing values associated with key.
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)//called by the handler_readiness
	//writing byte again to json
	w.Write(data)//Write writes the data to the connection as part of an HTTP response.
}