package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokecache"
	"io"
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
	Cache       *pokecache.Cache
}

func GetLocations(url string, config *Config) ([]Location, error) {
	cached, found := config.Cache.Get(url)
	if found {
		response, err := unmarshalAndUpdateConfig(cached, config)
		if err != nil {
			return nil, err
		}
		return response.Results, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	config.Cache.Add(url, body)

	response, err := unmarshalAndUpdateConfig(body, config)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

func unmarshalAndUpdateConfig(data []byte, config *Config) (ApiResponse, error) {
	var response ApiResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return ApiResponse{}, err
	}
	config.NextURL = response.Next
	config.PreviousURL = response.Previous
	return response, nil
}
