package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area/?limit=20&offset=0"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locations := Locations{}
		if err := json.Unmarshal(val, &locations); err != nil {
			return Locations{}, fmt.Errorf("error reading locations from cache: %v", err)
		}
		return locations, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, fmt.Errorf("error fetching map data: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, fmt.Errorf("error fetching map data: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Locations{}, fmt.Errorf("error: recived status code %d from server", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("error reading response body: %v", err)
	}
	var mapData Locations
	if err := json.Unmarshal(body, &mapData); err != nil {
		return Locations{}, fmt.Errorf("error parsing JSON response: %v", err)
	}
	c.cache.Add(url, body)
	return mapData, nil
}
