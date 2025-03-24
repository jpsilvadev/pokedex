package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	// Fallback to baseURL if pageURL is not provided
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check if URL is cached
	if cachedData, exists := c.cache.Get(url); exists {
		// Use cache
		var locationAreas LocationAreasResponse
		err := json.Unmarshal(cachedData, &locationAreas)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreas, nil
	}

	// if not in cache, make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	var locationAreas LocationAreasResponse
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	// add entry to cache
	c.cache.Add(url, data)

	return locationAreas, nil
}

// GetListOfPokemonInLocation retrieves a list of Pokémon available in a specified location area by name.
// It first checks the cache for the data and makes an HTTP request to the API if not cached.
// Returns the parsed PokemonInLocationResponse and an error if any occurs during the process.
func (c *Client) GetListOfPokemonInLocation(name string) (PokemonInLocationResponse, error) {
	locationURL := baseURL + "location-area" + "/" + name

	// Check if data is cached
	if cachedData, exists := c.cache.Get(locationURL); exists {
		var pokemonInLocation PokemonInLocationResponse
		err := json.Unmarshal(cachedData, &pokemonInLocation)
		if err != nil {
			return PokemonInLocationResponse{}, err
		}
		return pokemonInLocation, nil
	}

	// if not in cache, make the request
	req, err := http.NewRequest("GET", locationURL, nil)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}

	var pokemonInLocation PokemonInLocationResponse
	err = json.Unmarshal(data, &pokemonInLocation)
	if err != nil {
		return PokemonInLocationResponse{}, err
	}

	c.cache.Add(locationURL, data)
	return pokemonInLocation, nil
}
