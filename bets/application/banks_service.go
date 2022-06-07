package application

import (
	"my-bets/bets/domain/bank"
	"time"

	"github.com/google/uuid"
)

type BanksService struct {
	BankRepository bank.IBanksRepository
}

func NewBankService(bankRepository bank.IBanksRepository) *BanksService {
	return &BanksService{
		BankRepository: bankRepository,
	}
}

func (b *BanksService) CreateABank(initialValue float64) (bank.Bank, error) {
	bank := bank.NewBank(uuid.NewString(), initialValue, time.Now())

	return *bank, b.BankRepository.CreateABank(*bank)
}

func (b *BanksService) GetABank(id string) (bank.Bank, error) {
	return b.BankRepository.GetABank(id)
}
