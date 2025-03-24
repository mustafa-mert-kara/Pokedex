package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Type_name struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Stats      []Stat `json:"stats"`
	Types      []Type `json:"types"`
}

func (c *Client) GetPokemon(pokeName *string) (Pokemon, error) {
	if pokeName == nil || *pokeName == "" {
		return Pokemon{}, fmt.Errorf("Pokemon Name is Missing")
	}

	url := baseURL + "/pokemon/" + *pokeName

	cachedData, ok := c.Cache.Get(url)
	if !ok {
		fmt.Println("Not In The Cache", url)
		dat, err := createandDo(c, url)
		if err != nil {
			return Pokemon{}, err
		}
		c.Cache.Add(url, dat)
		cachedData = dat
	} else {
		fmt.Println("Providing From The Cache", url)
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(cachedData, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
