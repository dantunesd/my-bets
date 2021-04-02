package infrastructure

import "fmt"

type InMemoryCacheAdapter struct {
	DataCached map[string]interface{}
}

func NewInMemoryCacheAdapter() *InMemoryCacheAdapter {
	return &InMemoryCacheAdapter{
		DataCached: make(map[string]interface{}),
	}
}

func (i *InMemoryCacheAdapter) Add(key string, content interface{}) error {
	fmt.Println("creating in cache")
	i.DataCached[key] = content
	return nil
}

func (i *InMemoryCacheAdapter) Get(key string) (interface{}, bool) {
	fmt.Println("getting from cache")
	content, exists := i.DataCached[key]
	return content, exists
}

func (i *InMemoryCacheAdapter) Update(key string, content interface{}) error {
	fmt.Println("updating in db")
	i.DataCached[key] = content
	return nil
}
