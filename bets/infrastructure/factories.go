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

func BankRepositoryFactory() application.IBanksRepository {
	return NewBankCacheDecorator(
		&BankRepository{
			Database: DatabaseFactory("my-bets", "banks"),
		},
	)
}

func BetRepositoryFactory() application.IBetsRepository {
	return &BetRepository{
		Database: DatabaseFactory("my-bets", "bets"),
	}
}

func BankServiceFactory(bankRepository application.IBanksRepository) *application.BanksService {
	return application.NewBankService(bankRepository)
}

func BetServiceFactory(bankRepository application.IBanksRepository) *application.BetsService {
	return application.NewBetsService(
		&domain.PlaceABetService{},
		bankRepository,
		BetRepositoryFactory(),
	)
}

func HandlersFactory() http.Handler {
	bankRepository := BankRepositoryFactory()

	return presentation.HandlersFactory(
		BankServiceFactory(bankRepository),
		BetServiceFactory(bankRepository),
	)
}
