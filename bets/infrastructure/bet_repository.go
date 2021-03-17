package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BetRepository struct {
	Client *mongo.Client
}

func (b *BetRepository) Create(content interface{}) error {
	collection := b.Client.Database("my-bets").Collection("bets")

	_, err := collection.InsertOne(context.TODO(), content)
	return err
}

func (b *BetRepository) Get(id string, output interface{}) error {
	collection := b.Client.Database("my-bets").Collection("bets")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	return collection.FindOne(context.TODO(), filter).Decode(output)
}

func (b *BetRepository) Update(id string, content interface{}) error {
	return nil
}

func (b *BetRepository) Delete(id string) error {
	collection := b.Client.Database("my-bets").Collection("bets")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
