package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Config, arg string) error {
	url := pokeapi.BaseURL + arg
	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	pokemons, err := pokeapi.GetPokemons(url, config)
	if err != nil {
		return fmt.Errorf("Error fetching pokémons: %s\n", err)
	}
	for _, p := range pokemons {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
