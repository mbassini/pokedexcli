package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
	"os"
)

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
