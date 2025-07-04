package main

import "fmt"

func commandHelp(conf *config, args ...string) error {
	if args[0] == "" {
		fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
		for _, cmd := range getCommands() {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
	} else {
		if val, ok := getCommands()[args[0]]; ok {
			fmt.Printf("\nWelcome to the Pokedex!\nHelp page for %s command:\n\n", val.name)
			fmt.Printf("%s: %s\n", val.name, val.description)
		}
	}
	fmt.Println()
	return nil
}
