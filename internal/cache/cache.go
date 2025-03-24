package cache

import (
	"sync"
	"time"
)

func NewCache(dur time.Duration) *Cache {
	nCache := Cache{
		Data:  make(map[string]cacheEntry, 0),
		Mutex: &sync.Mutex{},
	}
	nCache.Ticker = nCache.reapLoop(dur)

	return &nCache

}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	c.Data[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
	c.Mutex.Unlock()

}

func (c *Cache) Get(key string) ([]byte, bool) {
	if data, ok := c.Data[key]; ok {
		return data.Val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(duration time.Duration) *time.Ticker {

	ticker := time.NewTicker(duration)
	go func(ticker *time.Ticker) {
		for {
			<-ticker.C
			for k, v := range c.Data {
				if time.Since(v.CreatedAt) > duration {
					c.Mutex.Lock()
					delete(c.Data, k)
					c.Mutex.Unlock()
				}
			}
		}

	}(ticker)
	return ticker

}
