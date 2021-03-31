package infrastructure

import (
	"my-bets/bets/application"
	"my-bets/bets/domain"
)

type BankCacheDecorator struct {
	Bank            map[string]domain.Bank
	BanksRepository application.IBanksRepository
}

func NewBankCacheDecorator(banksRepository application.IBanksRepository) *BankCacheDecorator {
	return &BankCacheDecorator{
		Bank:            make(map[string]domain.Bank),
		BanksRepository: banksRepository,
	}
}

func (b *BankCacheDecorator) CreateABank(bank domain.Bank) error {
	if err := b.BanksRepository.CreateABank((bank)); err != nil {
		return err
	}

	b.Bank[bank.ID] = bank
	return nil
}

func (b *BankCacheDecorator) GetABank(id string) (domain.Bank, error) {
	if cachedBank, existsInCache := b.Bank[id]; existsInCache {
		return cachedBank, nil
	}

	storedBank, err := b.BanksRepository.GetABank(id)
	if err != nil {
		return storedBank, err
	}

	b.Bank[storedBank.ID] = storedBank

	return storedBank, nil
}

func (b *BankCacheDecorator) UpdateABank(bank domain.Bank) error {
	if err := b.BanksRepository.UpdateABank(bank); err != nil {
		return err
	}

	b.Bank[bank.ID] = bank
	return nil
}
