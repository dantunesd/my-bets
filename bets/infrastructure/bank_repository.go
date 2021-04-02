package infrastructure

import (
	"fmt"
	"my-bets/bets/domain"
)

const bankIDKey = "_id"

type BankRepository struct {
	Database IDatabase
}

func NewBankRepository(database IDatabase) *BankRepository {
	return &BankRepository{
		Database: database,
	}
}

func (b *BankRepository) CreateABank(bank domain.Bank) error {
	fmt.Println("creating in db")
	return b.Database.Create(bank)
}

func (b *BankRepository) GetABank(id string) (domain.Bank, error) {
	fmt.Println("getting from db")
	var bank domain.Bank
	return bank, b.Database.Get(id, bankIDKey, &bank)
}

func (b *BankRepository) UpdateABank(bank domain.Bank) error {
	fmt.Println("updating in db")
	return b.Database.Update(bank.ID, bankIDKey, bank)
}
