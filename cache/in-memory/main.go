package main

import (
	"fmt"
	"sync"
)

type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]string)}
}

func (c *MemoryCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *MemoryCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {
	cache := NewMemoryCache()
	cache.Set("user:1", "Tayron")
	cache.Set("feature:checkout", "enabled")

	for _, key := range []string{"user:1", "feature:checkout", "missing"} {
		value, ok := cache.Get(key)
		if !ok {
			fmt.Printf("%s -> cache miss\n", key)
			continue
		}
		fmt.Printf("%s -> %s\n", key, value)
	}
}
