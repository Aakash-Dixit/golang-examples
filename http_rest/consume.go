package main

import (
	"io/ioutil"

	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("responseData : " + string(responseData))
}
