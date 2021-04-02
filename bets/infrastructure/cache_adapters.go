package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type ICacheRepository interface {
	Add(key string, content interface{})
	Get(key string, output interface{}) bool
	Update(key string, content interface{})
}

//------------------------------------------------------------------------------
// In Volatile Memory Cache Adapter
//------------------------------------------------------------------------------
type InMemoryCacheAdapter struct {
	DataCached map[string][]byte
}

func NewInMemoryCacheAdapter() ICacheRepository {
	return &InMemoryCacheAdapter{
		DataCached: make(map[string][]byte),
	}
}

func (i *InMemoryCacheAdapter) Add(key string, content interface{}) {
	fmt.Println("creating in cache")

	data, _ := json.Marshal(content)
	i.DataCached[key] = data
}

func (i *InMemoryCacheAdapter) Get(key string, output interface{}) bool {
	fmt.Println("getting from cache")

	content, exists := i.DataCached[key]
	json.Unmarshal([]byte(content), output)
	return exists
}

func (i *InMemoryCacheAdapter) Update(key string, content interface{}) {
	fmt.Println("updating in cache")

	data, _ := json.Marshal(content)
	i.DataCached[key] = data
}

//------------------------------------------------------------------------------
// Redis Cache Adapter
//------------------------------------------------------------------------------
type RedisCacheAdapter struct {
	Client *redis.Client
}

func NewRedisCacheAdapter(client *redis.Client) ICacheRepository {
	return &RedisCacheAdapter{
		Client: client,
	}
}

func (r *RedisCacheAdapter) Add(key string, content interface{}) {
	fmt.Println("creating in cache")

	marshalledContet, _ := json.Marshal(content)
	r.Client.Set(context.TODO(), key, marshalledContet, 0)
}

func (r *RedisCacheAdapter) Get(key string, output interface{}) bool {
	fmt.Println("getting from cache")

	data, err := r.Client.Get(context.TODO(), key).Result()
	json.Unmarshal([]byte(data), output)
	return err == nil
}

func (r *RedisCacheAdapter) Update(key string, content interface{}) {

	fmt.Println("updating in cache")
	marshalledContet, _ := json.Marshal(content)
	r.Client.Set(context.TODO(), key, marshalledContet, 0)
}
