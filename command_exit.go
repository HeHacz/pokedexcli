package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... \nGoodbye!")
	os.Exit(0)
	return nil
}
