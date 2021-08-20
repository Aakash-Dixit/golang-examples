package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	token, err := generateJWT()
	if err != nil {
		log.Fatal("error while generating valid token : ", err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	req.Header.Set("Token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error while making http request : ", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error while reading response body : ", err.Error())
	}
	log.Println("Response Body : ", string(body))
}

func generateJWT() (string, error) {
	signingKey := []byte("my_sign_key")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Test Client"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
