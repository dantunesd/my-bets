package infrastructure

import "fmt"

type InMemoryCacheRepository struct {
	DataCached map[string]interface{}
}

func NewInMemoryCacheRepository() *InMemoryCacheRepository {
	return &InMemoryCacheRepository{
		DataCached: make(map[string]interface{}),
	}
}

func (i *InMemoryCacheRepository) Add(key string, content interface{}) error {
	fmt.Println("creating in cache")
	i.DataCached[key] = content
	return nil
}

func (i *InMemoryCacheRepository) Get(key string) (interface{}, bool) {
	fmt.Println("getting from cache")
	content, exists := i.DataCached[key]
	return content, exists
}

func (i *InMemoryCacheRepository) Update(key string, content interface{}) error {
	fmt.Println("updating in db")
	i.DataCached[key] = content
	return nil
}
