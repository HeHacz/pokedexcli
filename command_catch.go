package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, pokemonName ...string) error {
	if pokemonName[0] == "" {
		fmt.Println("You need to know what you want to catch. Go explore the world!")
	}
	pokemon, err := conf.pokeapiClient.PokemonInfo(pokemonName[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemon data: %v", err)
	}
	real_name := pokemon.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", real_name)
	seed := rand.Intn(pokemon.BaseExperience / 20)
	if seed+1 == pokemon.BaseExperience/20 {
		conf.pokedex[real_name] = pokemon
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", real_name)
	} else {
		fmt.Printf("%s escaped!\n", real_name)
	}
	return nil
}
