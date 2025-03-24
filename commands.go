package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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

func commandExplore(cfg *config, areaName ...string) error {
	if len(areaName) == 0 {
		return errors.New("you need to provide an area name to explore")
	}

	if len(areaName) > 1 {
		return errors.New("you can only provide one area name to explore")
	}

	locationArea := areaName[0]

	pokemonsInLocation, err := cfg.pokeapiClient.GetListOfPokemonInLocation(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v\n", locationArea)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonsInLocation.PokemonEncounters {
		fmt.Printf(" - %v\n", string(pokemon.Pokemon.Name))
	}
	return nil
}

func commandCatch(cfg *config, pokemonName ...string) error {
	if len(pokemonName) == 0 {
		return errors.New("you need to provide a pokemon name to catch")
	}
	if len(pokemonName) > 1 {
		return errors.New("you can only provide one pokemon name to catch")
	}

	pokemonIdentifier := pokemonName[0]

	pokemonData, err := cfg.pokeapiClient.GetPokemonData(pokemonIdentifier)
	if err != nil {
		// Pokemon does not exist
		fmt.Println("Pokemon not found")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonIdentifier)

	pokemonEXP := pokemonData.BaseExperience
	randomChanceToCatch := rand.Intn(pokemonEXP)

	if randomChanceToCatch > 40 {
		fmt.Printf("%v escaped!\n", pokemonIdentifier)
		return nil
	}

	fmt.Printf("%v caught!\n", pokemonIdentifier)
	cfg.pokedex.AddPokemon(pokemonIdentifier, pokemonData)
	return nil
}
