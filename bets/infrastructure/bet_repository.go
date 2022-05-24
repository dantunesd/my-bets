package infrastructure

import (
	"my-bets/bets/application"
	"my-bets/bets/domain"
)

const betIDKey = "_id"

type BetRepository struct {
	Database IDatabase
}

func NewBetRepository(database IDatabase) application.IBetsRepository {
	return &BetRepository{
		Database: database,
	}
}

func (b *BetRepository) CreateABet(bet domain.Bet) error {
	return b.Database.Create(bet)
}

func (b *BetRepository) GetABet(id string) (domain.Bet, error) {
	var bet domain.Bet
	return bet, b.Database.Get(id, betIDKey, &bet)
}

func (b *BetRepository) DeleteABet(id string) error {
	return b.Database.Delete(id, betIDKey)
}
