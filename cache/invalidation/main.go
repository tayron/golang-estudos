package main

import (
	"fmt"
	"sync"
)

type Database struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewDatabase() *Database {
	return &Database{data: map[string]string{"user:42": "plano-basic"}}
}

func (db *Database) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.data[key]
	return value, ok
}

func (db *Database) Update(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Invalidate(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

type Service struct {
	db    *Database
	cache *Cache
}

func NewService() *Service {
	return &Service{db: NewDatabase(), cache: NewCache()}
}

func (s *Service) GetPlan(userID string) string {
	if value, ok := s.cache.Get(userID); ok {
		fmt.Println("cache hit")
		return value
	}

	fmt.Println("cache miss")
	value, _ := s.db.Get(userID)
	s.cache.Set(userID, value)
	return value
}

func (s *Service) UpdatePlan(userID, newPlan string) {
	s.db.Update(userID, newPlan)
	s.cache.Invalidate(userID)
	fmt.Println("cache invalidado apos atualizacao")
}

func main() {
	service := NewService()

	fmt.Println("leitura 1:", service.GetPlan("user:42"))
	fmt.Println("leitura 2:", service.GetPlan("user:42"))

	service.UpdatePlan("user:42", "plano-premium")

	fmt.Println("leitura 3:", service.GetPlan("user:42"))
}
