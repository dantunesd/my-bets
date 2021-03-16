package application

import (
	"my-bets/bets/domain"
)

type BetsService struct {
	PbService      domain.IPlaceABet
	BankRepository IRepository
	BetRepository  IRepository
}

func NewBetsService(pbService domain.IPlaceABet, bankRepository, betRepository IRepository) *BetsService {
	return &BetsService{
		PbService:      pbService,
		BankRepository: bankRepository,
		BetRepository:  betRepository,
	}
}

func (b *BetsService) PlaceABet(bet domain.Bet) (domain.Bet, error) {
	var bank domain.Bank

	if gerr := b.BankRepository.Get(bet.BankID, &bank); gerr != nil {
		return bet, gerr
	}

	if perr := b.PbService.PlaceABet(bet, &bank); perr != nil {
		return bet, perr
	}

	if cerr := b.BetRepository.Create(bet); cerr != nil {
		return bet, cerr
	}

	if uerr := b.BankRepository.Update(bank.ID, bank); uerr != nil {
		return bet, uerr
	}

	return bet, nil
}
