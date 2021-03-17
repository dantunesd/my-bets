package application

import (
	"my-bets/bets/domain"
	"time"

	"github.com/google/uuid"
)

type BanksService struct {
	BankRepository IRepository
}

func NewBankService(bankRepository IRepository) *BanksService {
	return &BanksService{
		BankRepository: bankRepository,
	}
}

func (b *BanksService) CreateABank(initialValue float64) (domain.Bank, error) {
	bank := domain.InitializeANewBank(uuid.NewString(), initialValue, time.Now())

	return *bank, b.BankRepository.Create(bank)
}

func (b *BanksService) GetABank(id string) domain.Bank {
	var bank domain.Bank
	b.BankRepository.Get(id, &bank)
	return bank
}
