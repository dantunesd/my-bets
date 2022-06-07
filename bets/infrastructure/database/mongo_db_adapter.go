package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBAdapter struct {
	Client    *mongo.Client
	DBName    string
	TableName string
}

func NewMongoDBAdapter(client *mongo.Client, dbName, tableName string) IDatabase {
	return &MongoDBAdapter{
		Client:    client,
		DBName:    dbName,
		TableName: tableName,
	}
}

func (d *MongoDBAdapter) Create(content interface{}) error {
	collection := d.Client.Database(d.DBName).Collection(d.TableName)

	_, err := collection.InsertOne(context.TODO(), content)
	return err
}

func (d *MongoDBAdapter) Get(id, idFieldName string, output interface{}) error {
	collection := d.Client.Database(d.DBName).Collection(d.TableName)

	filter := bson.D{primitive.E{Key: idFieldName, Value: id}}

	return collection.FindOne(context.TODO(), filter).Decode(output)
}

func (d *MongoDBAdapter) Update(id string, idFieldName string, content interface{}) error {
	collection := d.Client.Database(d.DBName).Collection(d.TableName)

	filter := bson.D{primitive.E{Key: idFieldName, Value: id}}
	update := bson.M{
		"$set": content,
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *MongoDBAdapter) Delete(id, idFieldName string) error {
	collection := d.Client.Database(d.DBName).Collection(d.TableName)

	filter := bson.D{primitive.E{Key: idFieldName, Value: id}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
