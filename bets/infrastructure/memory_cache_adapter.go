package infrastructure

import (
	"encoding/json"
	"fmt"
)

type MemoryCacheAdapter struct {
	DataCached map[string][]byte
}

func NewMemoryCacheAdapter() ICache {
	return &MemoryCacheAdapter{
		DataCached: make(map[string][]byte),
	}
}

func (i *MemoryCacheAdapter) Set(key string, content interface{}) {
	fmt.Println("setting in cache")

	data, _ := json.Marshal(content)
	i.DataCached[key] = data
}

func (i *MemoryCacheAdapter) Get(key string, output interface{}) bool {
	fmt.Println("getting from cache")

	content, exists := i.DataCached[key]
	json.Unmarshal([]byte(content), output)
	return exists
}
