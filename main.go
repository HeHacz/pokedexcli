package main

import (
	"time"

	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
)

var conf = config{}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(10 * time.Second),
	}
	startRepl(cfg)
}
