package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type PokemonApiResponse struct {
	ID                int                `json:"id"`
	Location          Location           `json:"location"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func GetPokemons(url string, config *Config) ([]Pokemon, error) {
	cached, found := config.Cache.Get(url)
	if found {
		var response PokemonApiResponse
		err := json.Unmarshal(cached, &response)
		if err != nil {
			return nil, err
		}
		var pokemons []Pokemon
		for _, p := range response.PokemonEncounters {
			pokemons = append(pokemons, p.Pokemon)
		}
		return pokemons, nil
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
	var response PokemonApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	var pokemons []Pokemon
	for _, p := range response.PokemonEncounters {
		pokemons = append(pokemons, p.Pokemon)
	}
	return pokemons, nil
}

type PokemonDetails struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func TryToCatchPokemon(url string, config *Config) bool {
	res, err := http.Get(url)
	if err != nil {
		return false
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return false
	}
	if err != nil {
		return false
	}

	config.Cache.Add(url, body)
	var pokemon PokemonDetails
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return false
	}

	pokemonClass := getClass(pokemon.BaseExperience)

	num := rand.Intn(100)
	caught := false
	switch pokemonClass {
	case "weak":
		caught = num <= 95
	case "common":
		caught = num <= 75
	case "rare":
		caught = num <= 40
	case "legendary":
		caught = num <= 15
	}

	if caught {
		config.Pokedex[pokemon.Name] = pokemon
	}
	return caught
}

func getClass(experience int) string {
	switch true {
	case experience <= 50:
		return "weak"
	case experience <= 100:
		return "common"
	case experience <= 200:
		return "rare"
	case experience <= 400:
		return "legendary"
	}
	return ""
}
