package infrastructure

import (
	"my-bets/bets/application"
	"my-bets/bets/domain"
)

type BankRepositoryDecorator struct {
	CacheRepository ICacheRepository
	BanksRepository application.IBanksRepository
}

func NewBankRepositoryDecorator(banksRepository application.IBanksRepository, cacheRepository ICacheRepository) *BankRepositoryDecorator {
	return &BankRepositoryDecorator{
		CacheRepository: cacheRepository,
		BanksRepository: banksRepository,
	}
}

func (b *BankRepositoryDecorator) CreateABank(bank domain.Bank) error {
	if err := b.BanksRepository.CreateABank((bank)); err != nil {
		return err
	}

	return b.CacheRepository.Add(bank.ID, bank)
}

func (b *BankRepositoryDecorator) GetABank(id string) (domain.Bank, error) {
	cachedBank, existsInCache := b.CacheRepository.Get(id)

	if existsInCache {
		return cachedBank.(domain.Bank), nil
	}

	storedBank, err := b.BanksRepository.GetABank(id)
	if err != nil {
		return storedBank, err
	}

	return storedBank, b.CacheRepository.Add(storedBank.ID, storedBank)
}

func (b *BankRepositoryDecorator) UpdateABank(bank domain.Bank) error {
	if err := b.BanksRepository.UpdateABank(bank); err != nil {
		return err
	}

	return b.CacheRepository.Update(bank.ID, bank)
}
