package application

import (
	"my-bets/bets/domain"
	"time"

	"github.com/google/uuid"
)

type BetsService struct {
	PbService      IPlaceABetService
	BankRepository IRepository
	BetRepository  IRepository
}

func NewBetsService(pbService IPlaceABetService, bankRepository, betRepository IRepository) *BetsService {
	return &BetsService{
		PbService:      pbService,
		BankRepository: bankRepository,
		BetRepository:  betRepository,
	}
}

type BetDTO struct {
	ID        string    `bson:"_id" json:"id"`
	BankID    string    `json:"bank_id"`
	Market    string    `json:"market"`
	Event     string    `json:"event"`
	EventDate time.Time `json:"event_date"`
	Value     float64   `json:"value"`
	Result    float64   `json:"result"`
	Odd       float64   `json:"odd"`
	CreatedAt time.Time `json:"created_at"`
	Free      bool      `json:"free"`
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

func (b *BetsService) PlaceABet(pbd PlaceBetDTO) (*BetDTO, error) {
	bet := domain.Bet{
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

	var bank domain.Bank

	if gerr := b.BankRepository.Get(bet.BankID, &bank); gerr != nil {
		return nil, gerr
	}

	if perr := b.PbService.PlaceABet(bet, &bank); perr != nil {
		return nil, perr
	}

	if cerr := b.BetRepository.Create(bet); cerr != nil {
		return nil, cerr
	}

	if uerr := b.BankRepository.Update(bank.ID, bank); uerr != nil {
		return nil, uerr
	}

	betDTO := BetDTO(bet)

	return &betDTO, nil
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
