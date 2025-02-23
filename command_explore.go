package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Config, arg string) error {
	url := pokeapi.BaseURL + arg
	fmt.Println("URL:", url)
	pokemons, err := pokeapi.GetPokemons(url, config)
	if err != nil {
		return fmt.Errorf("Error fetching pokémons: %s\n", err)
	}
	for _, p := range pokemons {
		fmt.Println(p.Name)
	}
	return nil
}
