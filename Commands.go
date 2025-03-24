package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "get a list of 20 areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "get a list of previous 20 areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "inspect current pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMapf(cfg *config) error {

	nLocations, err := cfg.pokeClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		fmt.Println(err)
	}
	cfg.nextLocationsURL = nLocations.Next
	cfg.prevLocationsURL = nLocations.Previous
	for _, location := range nLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	nLocations, err := cfg.pokeClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		fmt.Println(err)
	}
	cfg.nextLocationsURL = nLocations.Next
	cfg.prevLocationsURL = nLocations.Previous
	for _, location := range nLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config) error {
	if len(cfg.arg) == 0 {
		return fmt.Errorf("no id or name given for exploration")
	}
	name := cfg.arg[0]
	pokemons, err := cfg.pokeClient.ListPokemon(name)
	if err != nil {
		fmt.Println(err)
	}

	for _, poke := range pokemons.PokemonEncounter {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.arg) == 0 {
		return fmt.Errorf("no id or name given for catching Pokemon")
	}
	pokemon, err := cfg.pokeClient.GetPokemon(&cfg.arg[0])
	if err != nil {
		return fmt.Errorf("cannot find pokemon. error: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	difficulty := 100*(pokemon.Experience-36)/(608-32) + 20
	if difficulty > 100 {
		difficulty = 100
	}
	throwChance := rand.Intn(100)
	fmt.Println("Catch Chance:", difficulty, "attempt:", throwChance)
	if throwChance > difficulty {
		fmt.Println(pokemon.Name, " was caught!")
		cfg.UserPokedex.pokemons[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, " escaped!")
	}

	return nil
}

func commandInspect(cfg *config) error {
	if len(cfg.arg) == 0 {
		return fmt.Errorf("no id or name given for inspecting Pokemon")
	}
	if pokemon, ok := cfg.UserPokedex.pokemons[cfg.arg[0]]; ok {

		printString := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		for _, v := range pokemon.Stats {
			printString += fmt.Sprintf("\t-%s: %d\n", v.Stat.Name, v.BaseStat)
		}
		printString += "Types:\n"
		for _, v := range pokemon.Types {
			printString += fmt.Sprintf("\t- %s\n", v.Type_name.Name)
		}
		fmt.Print(printString)
	} else {
		fmt.Print("User Has Not Caught This Pokemon")
	}
	return nil
}

func commandPokedex(cfg *config) error {

	pokemons := cfg.UserPokedex.pokemons
	fmt.Println("Your Pokedex:")
	for _, v := range pokemons {
		fmt.Printf(" - %s\n", v.Name)
	}

	return nil
}
