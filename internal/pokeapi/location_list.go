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

	return locations, nil
}
