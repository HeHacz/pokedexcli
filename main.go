package main

import (
	"time"

	"github.com/hehacz/pokedexcli/internal/pokeapi"
)

var conf = config{}

func main() {
	conf := &config{
		pokeapiClient: pokeapi.NewClient(10*time.Second, 5*time.Second),
	}
	startRepl(conf)
}
