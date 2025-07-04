package main

import (
	"fmt"
	"sort"
)

func commandPokedex(conf *config, args ...string) error {
	const fw = 18
	fmt.Println("Your Pokedex:")
	keys := make([]string, 0, len(conf.pokedex))
	for key := range conf.pokedex {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return conf.pokedex[keys[i]].ID < conf.pokedex[keys[j]].ID
	})
	for _, k := range keys {
		fmt.Printf("- %-*s\t#%d\n", fw, k, conf.pokedex[k].ID)
	}

	return nil
}
