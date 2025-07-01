package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationPokemonList(locationName string) (LocationPokemonList, error) {
	url := baseURL + "/location-area/" + locationName
	if val, ok := c.cache.Get(url); ok {
		pokemons := LocationPokemonList{}
		if err := json.Unmarshal(val, &pokemons); err != nil {
			return LocationPokemonList{}, fmt.Errorf("error reading pokemons from cache: %v", err)
		}
		return pokemons, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemonList{}, fmt.Errorf("error fetching pokemons data: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemonList{}, fmt.Errorf("error fetching pokemons data: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return LocationPokemonList{}, fmt.Errorf("error: recived status code %d from server", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemonList{}, fmt.Errorf("error reading response body: %v", err)
	}
	var pokemonsData LocationPokemonList
	if err := json.Unmarshal(body, &pokemonsData); err != nil {
		return LocationPokemonList{}, fmt.Errorf("error parsing JSON response: %v", err)
	}
	c.cache.Add(url, body)
	return pokemonsData, nil
}
