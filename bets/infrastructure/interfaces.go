package infrastructure

type ICacheRepository interface {
	Add(key string, content interface{}) error
	Get(key string) (interface{}, bool)
	Update(key string, content interface{}) error
}

type IDatabase interface {
	Create(content interface{}) error
	Get(id, idFieldName string, output interface{}) error
	Update(id string, idFieldName string, content interface{}) error
	Delete(id, idFieldName string) error
}
