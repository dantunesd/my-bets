package repository

import (
	"my-bets/bets/domain"
)

type BankRepositoryCacheProxy struct {
	CacheRepository ICache
	BanksRepository domain.IBanksRepository
}

func NewBankRepositoryCacheProxy(banksRepository domain.IBanksRepository, cacheRepository ICache) domain.IBanksRepository {
	return &BankRepositoryCacheProxy{
		CacheRepository: cacheRepository,
		BanksRepository: banksRepository,
	}
}

func (b *BankRepositoryCacheProxy) CreateABank(bank domain.Bank) error {
	if err := b.BanksRepository.CreateABank((bank)); err != nil {
		return err
	}

	b.CacheRepository.Set(bank.ID, bank)
	return nil
}

func (b *BankRepositoryCacheProxy) GetABank(id string) (domain.Bank, error) {
	cachedBank := domain.Bank{}
	existsInCache := b.CacheRepository.Get(id, &cachedBank)

	if existsInCache {
		return cachedBank, nil
	}

	storedBank, err := b.BanksRepository.GetABank(id)
	if err != nil {
		return storedBank, err
	}
	b.CacheRepository.Set(storedBank.ID, storedBank)
	return storedBank, nil
}

func (b *BankRepositoryCacheProxy) UpdateABank(bank domain.Bank) error {
	if err := b.BanksRepository.UpdateABank(bank); err != nil {
		return err
	}

	b.CacheRepository.Set(bank.ID, bank)
	return nil
}
