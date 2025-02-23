package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		inputCmd := words[0]

		command, err := getCommand(inputCmd)
		if err != nil {
			fmt.Println(err)
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println("Error executing command %s: %s", command, err)
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
	callback    func() error
}

func getCommand(cmd string) (cliCommand, error) {
	availableCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help",
			callback:    commandHelp,
		},
	}
	command, exists := availableCommands[cmd]
	if !exists {
		return cliCommand{}, fmt.Errorf("Unknown command: %s", cmd)
	}
	return command, nil
}
