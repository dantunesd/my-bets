package application

import (
	"my-bets/bets/domain"
	"time"

	"github.com/google/uuid"
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
	bet.ID = uuid.NewString()
	bet.CreatedAt = time.Now()

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

func (b *BetsService) UndoABet(ID string) error {
	var bet domain.Bet
	var bank domain.Bank

	if berr := b.BetRepository.Get(ID, &bet); berr != nil {
		return berr
	}

	if gerr := b.BankRepository.Get(bet.BankID, &bank); gerr != nil {
		return gerr
	}

	if perr := b.PbService.UndoABet(bet, &bank); perr != nil {
		return perr
	}

	if derr := b.BetRepository.Delete(bet.ID); derr != nil {
		return derr
	}

	if uerr := b.BankRepository.Update(bank.ID, bank); uerr != nil {
		return uerr
	}

	return nil
}
