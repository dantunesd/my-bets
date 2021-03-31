package infrastructure

import (
	"context"
	"my-bets/bets/application"
	"my-bets/bets/domain"
	"my-bets/bets/presentation"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClientFactory() *mongo.Client {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	return client
}

func DatabaseFactory(dBName, tableName string) *Database {
	return &Database{
		Client:    MongoClientFactory(),
		DBName:    dBName,
		TableName: tableName,
	}
}

func BankRepositoryFactory() *BankRepository {
	return &BankRepository{
		Database: DatabaseFactory("my-bets", "banks"),
	}
}

func BetRepositoryFactory() *BetRepository {
	return &BetRepository{
		Database: DatabaseFactory("my-bets", "bets"),
	}
}

func BankServiceFactory() *application.BanksService {
	return application.NewBankService(BankRepositoryFactory())
}

func BetServiceFactory() *application.BetsService {
	return application.NewBetsService(
		&domain.PlaceABetService{},
		BankRepositoryFactory(),
		BetRepositoryFactory(),
	)
}

func HandlersFactory() http.Handler {
	return presentation.HandlersFactory(
		BankServiceFactory(),
		BetServiceFactory(),
	)
}
