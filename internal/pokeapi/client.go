package pokeapi

import (
	"net/http"
	"time"

	"github.com/mustafa-mert-kara/Pokedex/internal/cache"
)

// Client -
type Client struct {
	httpClient http.Client
	Cache      cache.Cache
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// NewClient -
func NewClient(timeout time.Duration) Client {
	Cache := cache.NewCache(5 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: *Cache,
	}
}
