package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("error reading pokemon info from cache: %v", err)
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error fetching pokemon data: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error fetching pokemon data: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("error: recived status code %d from server", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response body: %v", err)
	}
	var pokemonData Pokemon
	if err := json.Unmarshal(body, &pokemonData); err != nil {
		return Pokemon{}, fmt.Errorf("error parsing JSON response: %v", err)
	}
	c.cache.Add(url, body)
	return pokemonData, nil
}
