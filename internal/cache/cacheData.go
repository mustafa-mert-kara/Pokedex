package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Data   map[string]cacheEntry
	Mutex  *sync.Mutex
	Ticker *time.Ticker
}
