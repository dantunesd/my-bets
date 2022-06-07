package bets

import (
	"errors"
	"my-bets/bets/domain/bank"
)

type IPlaceABetService interface {
	PlaceABet(bet Bet, bank *bank.Bank) error
	UndoABet(bet Bet, bank *bank.Bank) error
}

type PlaceABetService struct{}

func NewPlaceABetService() *PlaceABetService {
	return &PlaceABetService{}
}

func (p *PlaceABetService) PlaceABet(bet Bet, bank *bank.Bank) error {

	if !bet.IsValidResult() {
		return errors.New("the result is not valid")
	}

	if bet.Value > bank.CurrentValue {
		return errors.New("the bet value cannot be higher than the bank currentValue")
	}

	if bet.NeedApplyFreeBet() && bank.IsValidForFreeBets() {
		return nil
	}

	bank.UpdateBank(bet.Result, bet.CreatedAt)

	return nil
}

func (p *PlaceABetService) UndoABet(bet Bet, bank *bank.Bank) error {
	bank.UpdateBank(-bet.Result, bet.CreatedAt)

	return nil
}
