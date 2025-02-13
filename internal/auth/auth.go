package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an API key from the headers of an http request
//eg:
func GetAPIKey(headers http.Request) (string,error) {
	//headers has the http request
	//returning the value of the key Authorization or error if occurred
	val := headers.Header.Get("Authorization")
	
	//val is the value of the key Authorization
	if val == "" {//if no value found
		return "", errors.New("no Authentication info found")
	}
	
	// Authorization : APIkey {insert apikey here}
	vals := strings.Split(val, " ")//splitting the value of the key Authorization in terms of space
	if len(vals) !=2{
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey"{//if the first part of the value is not APIkey
		return "", errors.New("malformed first part of the auth header")
	}
	//returning the second part of the value
	return vals[1], nil
}