package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandPokedex(config *pokeapi.Config, arg string) error {
	for _, p := range config.Pokedex {
		fmt.Printf("  - %s", p.Name)
	}
	return nil
}
