package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type locationDetail struct {
	Id               int `json:"id"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocations(pageUrl *string) (pokemonMap, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	cachedData, ok := c.Cache.Get(url)
	if !ok {
		fmt.Println("Not In The Cache", url)
		dat, err := createandDo(c, url)
		if err != nil {
			return pokemonMap{}, err
		}
		c.Cache.Add(url, dat)
		cachedData = dat
	} else {
		fmt.Println("Providing From The Cache", url)
	}

	pokeMap := pokemonMap{}
	err := json.Unmarshal(cachedData, &pokeMap)
	if err != nil {
		return pokemonMap{}, err
	}
	return pokeMap, nil
}

func (c *Client) ListPokemon(name string) (locationDetail, error) {
	if name == "" {
		return locationDetail{}, fmt.Errorf("no name or id given")
	}
	url := baseURL + "/location-area/" + name
	cachedData, ok := c.Cache.Get(url)
	if !ok {
		fmt.Println("Not In The Cache", url)
		dat, err := createandDo(c, url)
		if err != nil {
			return locationDetail{}, err
		}
		c.Cache.Add(url, dat)
		cachedData = dat
	}
	locDetails := locationDetail{}
	err := json.Unmarshal(cachedData, &locDetails)
	if err != nil {
		return locationDetail{}, err
	}

	return locDetails, nil
}

func createandDo(c *Client, url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return dat, nil

}
