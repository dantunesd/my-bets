package main

import (
	"context"
	"my-bets/bets/application"
	"my-bets/bets/domain/bets"
	"my-bets/bets/infrastructure/cache"
	"my-bets/bets/infrastructure/database"
	"my-bets/bets/infrastructure/repository"
	"my-bets/bets/presentation"
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
			database.NewLoggableDBDecorator(
				database.NewMongoDBAdapter(mongoClient, "my-bets", "banks"),
			),
		),
		cache.NewRedisCacheAdapter(redisClient),
	)

	betsRepository := repository.NewBetRepository(
		database.NewLoggableDBDecorator(
			database.NewMongoDBAdapter(mongoClient, "my-bets", "bets"),
		),
	)

	placeBetsService := bets.NewPlaceABetService()
	banksService := application.NewBankService(banksRepository)
	betsService := application.NewBetsService(
		placeBetsService,
		banksRepository,
		betsRepository,
	)

	banksHandler := presentation.NewBanksHandler(banksService)
	betsHandler := presentation.NewBetsHandler(betsService)

	httpHandler := presentation.NewHandler(
		banksHandler,
		betsHandler,
	)

	http.ListenAndServe(":8080", httpHandler.Create())
}
