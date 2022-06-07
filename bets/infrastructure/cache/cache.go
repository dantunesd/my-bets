package cache

type ICache interface {
	Set(key string, content interface{})
	Get(key string, output interface{}) bool
}
