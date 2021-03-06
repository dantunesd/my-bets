package domain

import "errors"

type PlaceABetService struct{}

func (p *PlaceABetService) PlaceABet(bet Bet, bank *Bank) error {

	if !bet.IsValidResult() {
		return errors.New("the result is not valid")
	}

	if bet.Value > bank.CurrentValue {
		return errors.New("the bet value cannot be higher than the bank currentValue")
	}

	if bet.needApplyFreeBet() && bank.isValidForFreeBets() {
		return nil
	}

	bank.UpdateBank(bet.Result, bet.CreatedAt)

	return nil
}

func (p *PlaceABetService) UndoABet(bet Bet, bank *Bank) error {
	bank.UpdateBank(-bet.Result, bet.CreatedAt)

	return nil
}
