package application

import "my-bets/bets/domain"

type IRepository interface {
	Create(content interface{}) error
	Get(id string, output interface{}) error
	Update(id string, content interface{}) error
	Delete(id string) error
}

type IPlaceABetService interface {
	PlaceABet(bet domain.Bet, bank *domain.Bank) error
	UndoABet(bet domain.Bet, bank *domain.Bank) error
}
