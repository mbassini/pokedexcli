package main

import (
	"fmt"
	"os"

	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
