package domain

import "errors"

type PlaceABetService struct{}

func (p *PlaceABetService) PlaceABet(bet Bet, bank Bank) (Bank, error) {

	if !bet.IsValidResult() {
		return bank, errors.New("the result is not valid")
	}

	if bet.Value > bank.CurrentValue {
		return bank, errors.New("the bet value cannot be higher than the bank currentValue")
	}

	bank.UpdateBank(bet.Result, bet.CreatedAt)

	return bank, nil
}
