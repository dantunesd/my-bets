package repository

import (
	"my-bets/bets/domain/bank"
	"my-bets/bets/infrastructure/database"
)

const bankIDKey = "_id"

type BankRepository struct {
	Database database.IDatabase
}

func NewBankRepository(database database.IDatabase) bank.IBanksRepository {
	return &BankRepository{
		Database: database,
	}
}

func (b *BankRepository) CreateABank(bank bank.Bank) error {
	return b.Database.Create(bank)
}

func (b *BankRepository) GetABank(id string) (bank.Bank, error) {
	var bank bank.Bank
	return bank, b.Database.Get(id, bankIDKey, &bank)
}

func (b *BankRepository) UpdateABank(bank bank.Bank) error {
	return b.Database.Update(bank.ID, bankIDKey, bank)
}
