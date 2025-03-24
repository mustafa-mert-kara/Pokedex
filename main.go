package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mustafa-mert-kara/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type Pokedex struct {
	pokemons map[string]pokeapi.Pokemon
}

type config struct {
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	UserPokedex      *Pokedex
	arg              []string
}

func main() {
	cfg := &config{
		pokeClient: pokeapi.NewClient(5 * time.Second),
		UserPokedex: &Pokedex{
			pokemons: make(map[string]pokeapi.Pokemon),
		},
	}

	defer cfg.pokeClient.Cache.Ticker.Stop()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)
		if command, ok := getCommands()[cleaned[0]]; ok {
			cfg.arg = cleaned[1:]
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("Error: ", err)
			}

		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
}
