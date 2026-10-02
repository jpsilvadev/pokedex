package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("You must provide a pokemon name to <catch>")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// decaying curve with fixed k for estimating catch chance
	// min or max aren't known ahead of time
	// unless we want to pull all pokemon data one at a time
	// and figure it out
	const k = 100.0
	catchChance := k / (k + float64(pokemon.BaseExperience))

	if rand.Float64() < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the `inspect` command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
