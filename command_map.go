package main

import (
	"fmt"
)

func commandMap(conf *config) error {

	locations, err := conf.pokeapiClient.LocationList(conf.nextURL)
	if err != nil {
		return fmt.Errorf("error fetching map data: %v", err)
	}
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%v\n", locations[i]["name"])
	}
	if prev, ok := locations.Previous; ok {
		conf.prevURL = prev
	} else {
		return fmt.Errorf("you are on the first page, no previous locations available")
	}
	if next, ok := locations.Next; ok {
		conf.nextURL = next
	} else {
		return fmt.Errorf("you are on the last page, no next locations available")
	}
	return nil
}

func commandMapb(conf *config) error {
	locations, err := conf.pokeapiClient.LocationList(conf.prevURL)
	if err != nil {
		return fmt.Errorf("error fetching map data: %v", err)
	}
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%v\n", locations[i]["name"])
	}
	if prev, ok := locations.Previous; ok {
		conf.prevURL = prev
	} else {
		return fmt.Errorf("you are on the first page, no previous locations available")
	}
	if next, ok := locations.Next; ok {
		conf.nextURL = next
	} else {
		return fmt.Errorf("you are on the last page, no next locations available")
	}
	return nil
}
