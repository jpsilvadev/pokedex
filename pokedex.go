package main

import "github.com/jpsilvadev/pokedex/internal/pokeapi"

type Pokedex struct {
	Caught map[string]pokeapi.PokemonData
}

func (p *Pokedex) AddPokemon(name string, pokemonData pokeapi.PokemonData) {
	p.Caught[name] = pokemonData
}

func (p *Pokedex) IsPokemonCaught(name string) bool {
	_, exists := p.Caught[name]
	return exists
}
