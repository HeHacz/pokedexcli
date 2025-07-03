package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hehacz/pokedexcli/internal/pokeapi"
)

type config struct {
	nextURL       *string
	prevURL       *string
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v", err)
			}
			break
		} else {
			input := cleanInput(scanner.Text())
			if len(input) == 0 {
				fmt.Println("No input provided.")
				continue
			}
			if cmd, ok := getCommands()[input[0]]; ok {
				if len(input) > 1 {
					if err := cmd.callback(cfg, input[1]); err != nil {
						fmt.Printf("Error executing command '%s' : %v\n", cmd.name, err)
					}
				} else {
					if err := cmd.callback(cfg, ""); err != nil {
						fmt.Printf("Error executing command '%s' : %v\n", cmd.name, err)
					}
				}
			} else {
				fmt.Printf("Unknown command: %s\n", input)
			}
		}
	}
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	str := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(str)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "\tExit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "\tDisplay this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "\tDisplay the list of locations in the Pokemon world.\n\tEach use displays list of 20 locations.\n\tEach consecutive use displays the next 20 locattions.\n\tUse mapb to display list of 20 prverious locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "\tDisplay the list of locations in the Pokemon world.\n\tEach use displays list of 20 locations.\n\tEach consecutive use displays the previous 20 locations.\n\tUse map to display list of 20 next locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "\tDisplay list of pokemons found in area.\n\tProvide location name as a parameter.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "\tTry to catch Pokemon\n\tProvide pokemon name as a parameter.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "\tPrint information about catched pokemon\n\tProvide pokemon name as a parameter.",
			callback:    commandInspect,
		},
	}
}
