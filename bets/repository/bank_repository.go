package repository

import (
	"my-bets/bets/domain"
)

const bankIDKey = "_id"

type BankRepository struct {
	Database IDatabase
}

func NewBankRepository(database IDatabase) domain.IBanksRepository {
	return &BankRepository{
		Database: database,
	}
}

func (b *BankRepository) CreateABank(bank domain.Bank) error {
	return b.Database.Create(bank)
}

func (b *BankRepository) GetABank(id string) (domain.Bank, error) {
	var bank domain.Bank
	return bank, b.Database.Get(id, bankIDKey, &bank)
}

func (b *BankRepository) UpdateABank(bank domain.Bank) error {
	return b.Database.Update(bank.ID, bankIDKey, bank)
}
