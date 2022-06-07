package repository

import (
	"my-bets/bets/domain/bank"
	"my-bets/bets/infrastructure/cache"
)

type BankRepositoryCacheProxy struct {
	CacheRepository cache.ICache
	BanksRepository bank.IBanksRepository
}

func NewBankRepositoryCacheProxy(banksRepository bank.IBanksRepository, cacheRepository cache.ICache) bank.IBanksRepository {
	return &BankRepositoryCacheProxy{
		CacheRepository: cacheRepository,
		BanksRepository: banksRepository,
	}
}

func (b *BankRepositoryCacheProxy) CreateABank(bank bank.Bank) error {
	if err := b.BanksRepository.CreateABank((bank)); err != nil {
		return err
	}

	b.CacheRepository.Set(bank.ID, bank)
	return nil
}

func (b *BankRepositoryCacheProxy) GetABank(id string) (bank.Bank, error) {
	cachedBank := bank.Bank{}
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

func (b *BankRepositoryCacheProxy) UpdateABank(bank bank.Bank) error {
	if err := b.BanksRepository.UpdateABank(bank); err != nil {
		return err
	}

	b.CacheRepository.Set(bank.ID, bank)
	return nil
}
