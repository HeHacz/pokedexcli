package main

import (
	"fmt"
)

func commandExplore(conf *config, locationName ...string) error {
	if locationName[0] == "" {
		return fmt.Errorf("you can't catch anything if you don't start your journey")
	}
	pokemons, err := conf.pokeapiClient.LocationPokemonList(locationName[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemons list: %v", err)
	}
	fmt.Printf("Exploring %s\nPokemons to catch:\n", pokemons.Name)
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon.Name)
	}
	return nil
}
