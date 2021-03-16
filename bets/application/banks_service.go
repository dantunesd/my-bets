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

func (b *BanksService) CreateABank(initialValue float64) (string, error) {
	id := uuid.NewString()
	bank := domain.InitializeANewBank(id, initialValue, time.Now())

	return id, b.BankRepository.Create(bank)
}
