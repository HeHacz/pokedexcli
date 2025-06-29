package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... \nGoodbye!")
	os.Exit(0)
	return nil
}
