package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type DiskEntry struct {
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type DiskCache struct {
	basePath string
}

func NewDiskCache(basePath string) *DiskCache {
	return &DiskCache{basePath: basePath}
}

func (c *DiskCache) filePath(key string) string {
	return filepath.Join(c.basePath, key+".json")
}

func (c *DiskCache) Set(key, value string) error {
	if err := os.MkdirAll(c.basePath, 0o755); err != nil {
		return err
	}

	entry := DiskEntry{
		Value:     value,
		CreatedAt: time.Now(),
	}

	content, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.filePath(key), content, 0o644)
}

func (c *DiskCache) Get(key string) (DiskEntry, error) {
	content, err := os.ReadFile(c.filePath(key))
	if err != nil {
		return DiskEntry{}, err
	}

	var entry DiskEntry
	if err := json.Unmarshal(content, &entry); err != nil {
		return DiskEntry{}, err
	}

	return entry, nil
}

func main() {
	cache := NewDiskCache("./cache-data")

	if err := cache.Set("report-2026", "arquivo grande de relatorio"); err != nil {
		panic(err)
	}

	entry, err := cache.Get("report-2026")
	if err != nil {
		panic(err)
	}

	fmt.Println("valor:", entry.Value)
	fmt.Println("criado em:", entry.CreatedAt.Format(time.RFC3339))
}
