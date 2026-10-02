package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// look in cache first
	if val, ok := c.cache.Get(url); ok {
		locations := locationAreaResponse{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return locationAreaResponse{}, nil
		}
		return locations, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return locationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locationAreaResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}

	var locations locationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return locationAreaResponse{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}

func (c *Client) GetLocation(locationName string) (locationResponse, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		location := locationResponse{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return locationResponse{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return locationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locationResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse{}, err
	}

	var location locationResponse
	if err := json.Unmarshal(data, &location); err != nil {
		return locationResponse{}, err
	}

	c.cache.Add(url, data)
	return location, nil
}
