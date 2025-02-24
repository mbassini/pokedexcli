package main

import (
	"fmt"
	"github.com/mbassini/pokedexcli/internal/pokeapi"
)

func commandInspect(config *pokeapi.Config, arg string) error {
	pokemon, err := pokeapi.InspectPokemon(arg, config)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	return nil
}
