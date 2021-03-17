package main

import (
	"context"
	"my-bets/bets/application"
	"my-bets/bets/domain"
	"my-bets/bets/infrastructure"
	"my-bets/bets/presentation"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// @TODO move string to env var
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	bankRepository := infrastructure.BankRepository{Client: client}
	betRepository := infrastructure.BetRepository{Client: client}
	placeABetService := domain.PlaceABetService{}

	banksService := application.NewBankService(&bankRepository)
	betsService := application.NewBetsService(&placeABetService, &bankRepository, &betRepository)

	presentation.CreateAndStartServer(banksService, betsService)
}
