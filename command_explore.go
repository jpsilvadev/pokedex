package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("You must provide a location name to <explore>")
	}
	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)

	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounters := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounters.Pokemon.Name)
	}

	return nil
}
