package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var signingKey = []byte("my_sign_key")

	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method for token does not match")
			}
			return signingKey, nil
		})

		if err != nil {
			log.Println("Error occured while parsing and validating token : ", err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		if token.Valid {
			log.Println("Client is authorized, returning response")
			fmt.Fprintf(w, "Hello World")
		}
	} else {
		fmt.Fprintf(w, "Client is not authorized to preform the request")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
