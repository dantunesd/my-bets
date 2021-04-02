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

	b.CacheRepository.Add(bank.ID, bank)
	return nil
}

func (b *BankRepositoryDecorator) GetABank(id string) (domain.Bank, error) {
	cachedBank := domain.Bank{}
	existsInCache := b.CacheRepository.Get(id, &cachedBank)

	if existsInCache {
		return cachedBank, nil
	}

	storedBank, err := b.BanksRepository.GetABank(id)
	if err != nil {
		return storedBank, err
	}
	b.CacheRepository.Add(storedBank.ID, storedBank)
	return storedBank, nil
}

func (b *BankRepositoryDecorator) UpdateABank(bank domain.Bank) error {
	if err := b.BanksRepository.UpdateABank(bank); err != nil {
		return err
	}

	b.CacheRepository.Update(bank.ID, bank)
	return nil
}
