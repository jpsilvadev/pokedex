package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jpsilvadev/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	NextUrl       *string
	PreviousURL   *string
	pokedex       Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandProvided := words[0]
		command, exists := getCommands()[commandProvided]
		if exists {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next page of pokemon locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page of pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Display the available pokemon's in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Throw a pokeball at a pokemon",
			callback:    commandCatch,
		},
	}
}
