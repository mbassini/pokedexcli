package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandCatch(config *pokeapi.Config, arg string) error {
	_, alreadyCaught := config.Pokedex[arg]
	if alreadyCaught {
		fmt.Printf("%s already caught!\n", arg)
		return nil
	}

	url := pokeapi.BaseURL + "pokemon/" + arg
	fmt.Printf("Throwing a pockebal at %s...\n", arg)
	caught := pokeapi.TryToCatchPokemon(url, config)
	if !caught {
		fmt.Printf("%s escaped!\n", arg)
		return nil
	}
	fmt.Printf("%s was caught!\n", arg)

	return nil
}
