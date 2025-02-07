package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	fmt.Println("hello")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the enviornment")
	}
}