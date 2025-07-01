package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	locations, err := conf.pokeapiClient.LocationList(conf.nextURL)
	if err != nil {
		return fmt.Errorf("error fetching map data: %v", err)
	}
	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)

	}
	conf.prevURL = locations.Previous
	conf.nextURL = locations.Next
	if conf.nextURL == nil {
		return fmt.Errorf("you are on the last page, there are no more locations available")
	}
	return nil
}

func commandMapb(conf *config) error {
	if conf.prevURL == nil {
		return fmt.Errorf("you are on the first page, no previous locations available")
	}
	locations, err := conf.pokeapiClient.LocationList(conf.prevURL)
	if err != nil {
		return fmt.Errorf("error fetching map data: %v", err)
	}
	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	conf.prevURL = locations.Previous
	conf.nextURL = locations.Next

	return nil
}
