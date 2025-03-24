package main

import (
	"time"

	"github.com/jpsilvadev/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: Pokedex{
			Caught: make(map[string]pokeapi.PokemonData),
		},
	}
	startRepl(cfg)
}
