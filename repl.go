package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mbassini/pokedexcli/internal/pokeapi"
	"github.com/mbassini/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{}
	config.Cache = pokecache.NewCache(10 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		inputCmd := words[0]

		command, exists := getCommands()[inputCmd]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			if command.name == "explore" {
				err := command.callback(config, words[1])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(config, "")
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Displays the Pokémon located in the area",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandBMap,
		},
		"help": {
			name:        "help",
			description: "Show this help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
