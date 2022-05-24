package infrastructure

import (
	"fmt"
)

type LoggableCacheDecorator struct {
	Cache ICache
}

func NewLoggableCacheDecorator(cache ICache) ICache {
	return &LoggableCacheDecorator{
		Cache: cache,
	}
}

func (l *LoggableCacheDecorator) Set(key string, content interface{}) {
	fmt.Println("setting in cache", key)
	l.Cache.Set(key, content)

}

func (l *LoggableCacheDecorator) Get(key string, output interface{}) bool {
	fmt.Println("getting from cache", key)
	return l.Cache.Get(key, output)
}
