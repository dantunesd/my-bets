package infrastructure

import "fmt"

type LoggableDBDecorator struct {
	Database IDatabase
}

func NewLoggableDBDecorator(database IDatabase) IDatabase {
	return &LoggableDBDecorator{
		Database: database,
	}
}

func (d *LoggableDBDecorator) Create(content interface{}) error {
	fmt.Println("creating in DB", content)
	return d.Database.Create(content)
}

func (d *LoggableDBDecorator) Get(id, idFieldName string, output interface{}) error {
	fmt.Println("getting from DB", id)
	return d.Database.Get(id, idFieldName, output)
}

func (d *LoggableDBDecorator) Update(id string, idFieldName string, content interface{}) error {
	fmt.Println("updating in DB", id)
	return d.Database.Update(id, idFieldName, content)
}

func (d *LoggableDBDecorator) Delete(id, idFieldName string) error {
	fmt.Println("deleting in DB", id)
	return d.Database.Delete(id, idFieldName)
}
