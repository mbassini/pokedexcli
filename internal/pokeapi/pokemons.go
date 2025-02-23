package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonApiResponse struct {
	ID                int       `json:"id"`
	Location          Location  `json:"location"`
	PokemonEncounters []Pokemon `json:"pokemon"`
}

func GetPokemons(url string, config *Config) ([]Pokemon, error) {
	fmt.Println("Hello Starting GeFetch")
	cached, found := config.Cache.Get(url)
	if found {
		var response PokemonApiResponse
		err := json.Unmarshal(cached, &response)
		if err != nil {
			return nil, err
		}
		return response.PokemonEncounters, nil
	}
	fmt.Println("Hello Starting GeFetch")
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

	var response PokemonApiResponse
	err = json.Unmarshal(cached, &response)
	if err != nil {
		return nil, err
	}

	return response.PokemonEncounters, nil
}
