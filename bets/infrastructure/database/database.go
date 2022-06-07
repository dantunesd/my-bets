package database

type IDatabase interface {
	Create(content interface{}) error
	Get(id, idFieldName string, output interface{}) error
	Update(id string, idFieldName string, content interface{}) error
	Delete(id, idFieldName string) error
}
