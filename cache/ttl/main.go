package main

import (
	"fmt"
	"sync"
	"time"
)

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

type TTLCache struct {
	mu   sync.RWMutex
	data map[string]ttlEntry
}

func NewTTLCache() *TTLCache {
	return &TTLCache{data: make(map[string]ttlEntry)}
}

func (c *TTLCache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = ttlEntry{value: value, expiresAt: time.Now().Add(ttl)}
}

func (c *TTLCache) Get(key string) (string, bool) {
	c.mu.RLock()
	entry, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return "", false
	}

	if time.Now().After(entry.expiresAt) {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return "", false
	}

	return entry.value, true
}

func main() {
	cache := NewTTLCache()
	cache.Set("token:checkout", "abc123", 2*time.Second)

	value, ok := cache.Get("token:checkout")
	fmt.Println("antes de expirar:", value, ok)

	time.Sleep(3 * time.Second)

	value, ok = cache.Get("token:checkout")
	fmt.Println("depois de expirar:", value, ok)
}
