package main

import (
	"errors"
	"github.com/jpsilvadev/pokedex/internal/pokeapi"
)

type Pokedex struct {
	Caught map[string]pokeapi.PokemonData
}

type PokemonStat struct {
	Name  string
	Value int
}

type PokemonInfo struct {
	Name   string
	Height int
	Weight int
	Stats  []PokemonStat
	Types  []string
}

func (p *Pokedex) AddPokemon(name string, pokemonData pokeapi.PokemonData) {
	p.Caught[name] = pokemonData
}

func (p *Pokedex) IsPokemonCaught(name string) bool {
	_, exists := p.Caught[name]
	return exists
}

func (p *Pokedex) InspectPokemon(name string) (PokemonInfo, error) {
	if !p.IsPokemonCaught(name) {
		return PokemonInfo{}, errors.New("you have not caught that pokemon")
	}

	pokemonData := p.Caught[name]
	return newPokemonInfo(name, pokemonData), nil
}

func newPokemonInfo(name string, data pokeapi.PokemonData) PokemonInfo {
	var stats []PokemonStat
	for _, s := range data.Stats {
		stats = append(stats, PokemonStat{
			Name:  s.Stat.Name,
			Value: s.BaseStat,
		})
	}

	var types []string
	for _, t := range data.Types {
		types = append(types, t.Type.Name)
	}

	return PokemonInfo{
		Name:   name,
		Height: data.Height,
		Weight: data.Weight,
		Stats:  stats,
		Types:  types,
	}
}

func (p *Pokedex) ListCaughtPokemon() []string {
	var names []string
	for name := range p.Caught {
		names = append(names, name)
	}
	return names
}
