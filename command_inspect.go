package main

import (
	"fmt"
)

func commandInspect(conf *config, pokemonName ...string) error {
	if pokemonName[0] == "" {
		return fmt.Errorf("you need provide pokemon name")
	}
	if pokemon, ok := conf.pokedex[pokemonName[0]]; !ok {
		fmt.Printf("you have to catch %s first\n", pokemonName[0])
	} else {
		fmt.Printf("Name: %s #%d\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.ID, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, ty := range pokemon.Types {
			fmt.Printf("\t-%s\n", ty.Type.Name)
		}
	}

	return nil
}
