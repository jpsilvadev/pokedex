package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {

	// Fallback to baseURL if pageURL is not provided
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

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

	return locationAreas, nil
}
