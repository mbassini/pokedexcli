package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	url := pokeapi.BaseURL
	if config.NextURL != nil {
		url = *config.NextURL
	}
	locations, err := pokeapi.GetLocations(url, config)
	if err != nil {
		return fmt.Errorf("Error fetching maps: %s\n", err)
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}

func commandBMap(config *pokeapi.Config) error {
	if config.PreviousURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	url := *config.PreviousURL
	locations, err := pokeapi.GetLocations(url, config)
	if err != nil {
		return fmt.Errorf("Error fetching maps: %s\n", err)
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}
