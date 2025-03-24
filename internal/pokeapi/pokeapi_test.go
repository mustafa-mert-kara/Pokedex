package pokeapi_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mustafa-mert-kara/Pokedex/internal/pokeapi"
)

func TestResponse(t *testing.T) {
	client := pokeapi.NewClient(5 * time.Second)
	defer client.Cache.Ticker.Stop()
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "https://pokeapi.co/api/v2/location-area",
			expected: "canalave-city-area",
		},
		{
			input:    "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			expected: "mt-coronet-1f-route-216",
		},
	}
	for i, v := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			mapLocation, err := client.ListLocations(&v.input)
			if err != nil {
				t.Errorf("Couldn't Complete Request %v", err)
			}
			if mapLocation.Results[0].Name != v.expected {
				t.Errorf("Unexpected Results. Expected %v Found: %v", v.expected, mapLocation.Results[0].Name)
			}
		})
	}
}
