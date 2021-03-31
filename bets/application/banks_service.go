package application

import (
	"my-bets/bets/domain"
	"time"

	"github.com/google/uuid"
)

type BanksService struct {
	BankRepository IBanksRepository
}

func NewBankService(bankRepository IBanksRepository) *BanksService {
	return &BanksService{
		BankRepository: bankRepository,
	}
}

func (b *BanksService) CreateABank(initialValue float64) (domain.Bank, error) {
	bank := domain.NewBank(uuid.NewString(), initialValue, time.Now())

	return *bank, b.BankRepository.CreateABank(*bank)
}

func (b *BanksService) GetABank(id string) (domain.Bank, error) {
	return b.BankRepository.GetABank(id)
}
