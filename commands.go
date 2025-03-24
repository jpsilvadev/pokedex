package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.NextUrl)
	if err != nil {
		return err
	}

	cfg.NextUrl = locationAreas.Next
	cfg.PreviousURL = locationAreas.Previous

	for _, locArea := range locationAreas.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.PreviousURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.PreviousURL)
	if err != nil {
		return err
	}

	cfg.NextUrl = locationAreas.Next
	cfg.PreviousURL = locationAreas.Previous

	for _, locArea := range locationAreas.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}
