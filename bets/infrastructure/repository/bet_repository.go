package repository

import (
	"my-bets/bets/domain/bets"
	"my-bets/bets/infrastructure/database"
)

const betIDKey = "_id"

type BetRepository struct {
	Database database.IDatabase
}

func NewBetRepository(database database.IDatabase) bets.IBetsRepository {
	return &BetRepository{
		Database: database,
	}
}

func (b *BetRepository) CreateABet(bet bets.Bet) error {
	return b.Database.Create(bet)
}

func (b *BetRepository) GetABet(id string) (bets.Bet, error) {
	var bet bets.Bet
	return bet, b.Database.Get(id, betIDKey, &bet)
}

func (b *BetRepository) DeleteABet(id string) error {
	return b.Database.Delete(id, betIDKey)
}
