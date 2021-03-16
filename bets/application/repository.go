package application

type IRepository interface {
	Create(content interface{}) error
	Get(id string, output interface{}) error
	Update(id string, content interface{}) error
}
