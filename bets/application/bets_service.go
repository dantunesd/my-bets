package application

import (
	"my-bets/bets/domain/bank"
	"my-bets/bets/domain/bets"
	"time"

	"github.com/google/uuid"
)

type BetsService struct {
	PbService      bets.IPlaceABetService
	BankRepository bank.IBanksRepository
	BetRepository  bets.IBetsRepository
}

func NewBetsService(pbService bets.IPlaceABetService, bankRepository bank.IBanksRepository, betRepository bets.IBetsRepository) *BetsService {
	return &BetsService{
		PbService:      pbService,
		BankRepository: bankRepository,
		BetRepository:  betRepository,
	}
}

type PlaceBetDTO struct {
	BankID    string    `json:"bank_id"`
	Market    string    `json:"market"`
	Event     string    `json:"event"`
	EventDate time.Time `json:"event_date"`
	Value     float64   `json:"value"`
	Result    float64   `json:"result"`
	Odd       float64   `json:"odd"`
	Free      bool      `json:"free"`
}

func (b *BetsService) PlaceABet(pbd PlaceBetDTO) (*bets.Bet, error) {
	bet := bets.Bet{
		ID:        uuid.NewString(),
		BankID:    pbd.BankID,
		Market:    pbd.Market,
		Event:     pbd.Event,
		EventDate: pbd.EventDate,
		Value:     pbd.Value,
		Result:    pbd.Result,
		Odd:       pbd.Odd,
		Free:      pbd.Free,
		CreatedAt: time.Now(),
	}

	bank, gerr := b.BankRepository.GetABank(bet.BankID)
	if gerr != nil {
		return nil, gerr
	}

	if perr := b.PbService.PlaceABet(bet, &bank); perr != nil {
		return nil, perr
	}

	if cerr := b.BetRepository.CreateABet(bet); cerr != nil {
		return nil, cerr
	}

	if uerr := b.BankRepository.UpdateABank(bank); uerr != nil {
		return nil, uerr
	}

	return &bet, nil
}

func (b *BetsService) UndoABet(ID string) error {
	bet, berr := b.BetRepository.GetABet(ID)
	if berr != nil {
		return berr
	}

	bank, gerr := b.BankRepository.GetABank(bet.BankID)
	if gerr != nil {
		return gerr
	}

	if perr := b.PbService.UndoABet(bet, &bank); perr != nil {
		return perr
	}

	if derr := b.BetRepository.DeleteABet(bet.ID); derr != nil {
		return derr
	}

	if uerr := b.BankRepository.UpdateABank(bank); uerr != nil {
		return uerr
	}

	return nil
}
