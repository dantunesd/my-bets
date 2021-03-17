package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BankRepository struct {
	Client *mongo.Client
}

func (b *BankRepository) Create(content interface{}) error {
	collection := b.Client.Database("my-bets").Collection("banks")

	_, err := collection.InsertOne(context.TODO(), content)
	return err
}

func (b *BankRepository) Get(id string, output interface{}) error {
	collection := b.Client.Database("my-bets").Collection("banks")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	return collection.FindOne(context.TODO(), filter).Decode(output)
}

func (b *BankRepository) Update(id string, content interface{}) error {
	collection := b.Client.Database("my-bets").Collection("banks")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.M{
		"$set": content,
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (b *BankRepository) Delete(id string) error {
	collection := b.Client.Database("my-bets").Collection("banks")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
