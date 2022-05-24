package infrastructure

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisCacheAdapter struct {
	Client *redis.Client
}

func NewRedisCacheAdapter(client *redis.Client) ICache {
	return &RedisCacheAdapter{
		Client: client,
	}
}

func (r *RedisCacheAdapter) Set(key string, content interface{}) {
	marshalledContet, _ := json.Marshal(content)
	r.Client.Set(context.TODO(), key, marshalledContet, 0)
}

func (r *RedisCacheAdapter) Get(key string, output interface{}) bool {
	data, err := r.Client.Get(context.TODO(), key).Result()
	json.Unmarshal([]byte(data), output)
	return err == nil
}
