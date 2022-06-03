package main

import (
	"context"
	"my-bets/bets/application"
	"my-bets/bets/domain"
	"my-bets/bets/infrastructure"
	"my-bets/bets/presentation"
	"my-bets/bets/repository"
	"net/http"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoClient, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})

	banksRepository := repository.NewBankRepositoryCacheProxy(
		repository.NewBankRepository(
			infrastructure.NewLoggableDBDecorator(
				infrastructure.NewMongoDBAdapter(mongoClient, "my-bets", "banks"),
			),
		),
		infrastructure.NewRedisCacheAdapter(redisClient),
	)

	betsRepository := repository.NewBetRepository(
		infrastructure.NewLoggableDBDecorator(
			infrastructure.NewMongoDBAdapter(mongoClient, "my-bets", "bets"),
		),
	)

	banksService := application.NewBankService(banksRepository)
	betsService := application.NewBetsService(domain.NewPlaceABetService(), banksRepository, betsRepository)

	router := presentation.NewHandler(banksService, betsService)

	http.ListenAndServe(":8080", router.Create())
}
