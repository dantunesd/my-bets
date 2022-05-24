package application

import "my-bets/bets/domain"

type IBanksRepository interface {
	CreateABank(domain.Bank) error
	GetABank(id string) (domain.Bank, error)
	UpdateABank(domain.Bank) error
}

type IBetsRepository interface {
	CreateABet(domain.Bet) error
	GetABet(id string) (domain.Bet, error)
	DeleteABet(id string) error
}

type IPlaceABetService interface {
	PlaceABet(bet domain.Bet, bank *domain.Bank) error
	UndoABet(bet domain.Bet, bank *domain.Bank) error
}
