package infrastructure

import (
	"context"
	"my-bets/bets/application"
	"my-bets/bets/domain"
	"my-bets/bets/presentation"
	"net/http"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClientFactory() *mongo.Client {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	return client
}

func RedisClientFactory() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func DatabaseFactory(dBName, tableName string) IDatabase {
	return NewLoggableDBDecorator(
		NewMongoDBAdapter(
			MongoClientFactory(),
			dBName,
			tableName,
		),
	)
}

func CacheFactory(selectedCache string) ICache {
	availableCaches := map[string]ICache{
		"redis":  NewRedisCacheAdapter(RedisClientFactory()),
		"memory": NewMemoryCacheAdapter(),
	}
	return NewLoggableCacheDecorator(
		availableCaches[selectedCache],
	)
}

func BankRepositoryFactory() application.IBanksRepository {
	return NewBankRepositoryDecorator(
		NewBankRepository(DatabaseFactory("my-bets", "banks")),
		CacheFactory("redis"),
	)
}

func BetRepositoryFactory() application.IBetsRepository {
	return NewBetRepository(DatabaseFactory("my-bets", "bets"))
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
