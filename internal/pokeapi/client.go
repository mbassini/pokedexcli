package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2/location-area/"

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ApiResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Config struct {
	NextURL     *string
	PreviousURL *string
}

func GetLocations(url string, config *Config) ([]Location, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %s\n", err)
	}

	config.NextURL = response.Next
	config.PreviousURL = response.Previous
	return response.Results, nil
}
