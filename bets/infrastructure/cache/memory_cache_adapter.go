package cache

import (
	"encoding/json"
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
	data, _ := json.Marshal(content)
	i.DataCached[key] = data
}

func (i *MemoryCacheAdapter) Get(key string, output interface{}) bool {
	content, exists := i.DataCached[key]
	json.Unmarshal([]byte(content), output)
	return exists
}
