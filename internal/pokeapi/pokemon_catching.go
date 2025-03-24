package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(name string) (PokemonData, error) {
	pokemonURL := baseURL + "pokemon" + "/" + name

	// check if data is cached
	if cachedData, exists := c.cache.Get(pokemonURL); exists {
		var pokemon PokemonData
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonData{}, err
		}
		return pokemon, nil
	}

	// if not in cache, make the request
	req, err := http.NewRequest("GET", pokemonURL, nil)
	if err != nil {
		return PokemonData{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	var pokemon PokemonData
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonData{}, err
	}
	c.cache.Add(pokemonURL, data)
	return pokemon, nil
}
