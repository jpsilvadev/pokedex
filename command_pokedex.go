package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range cfg.caughtPokemon {
		fmt.Println("  -", pokemonName)
	}
	return nil
}
