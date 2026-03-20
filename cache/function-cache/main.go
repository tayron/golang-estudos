package main

import (
	"fmt"
	"sync"
	"time"
)

type FunctionCache struct {
	mu   sync.RWMutex
	data map[int]int
}

func NewFunctionCache() *FunctionCache {
	return &FunctionCache{data: make(map[int]int)}
}

func (c *FunctionCache) Memoize(fn func(int) int) func(int) int {
	return func(input int) int {
		c.mu.RLock()
		cached, ok := c.data[input]
		c.mu.RUnlock()
		if ok {
			fmt.Println("resultado vindo do cache")
			return cached
		}

		fmt.Println("executando funcao cara")
		result := fn(input)

		c.mu.Lock()
		c.data[input] = result
		c.mu.Unlock()
		return result
	}
}

func expensiveSquare(n int) int {
	time.Sleep(800 * time.Millisecond)
	return n * n
}

func main() {
	cache := NewFunctionCache()
	memoized := cache.Memoize(expensiveSquare)

	start := time.Now()
	fmt.Println("primeira chamada:", memoized(12), "tempo:", time.Since(start))

	start = time.Now()
	fmt.Println("segunda chamada:", memoized(12), "tempo:", time.Since(start))
}
