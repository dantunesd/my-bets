package domain

import "time"

type Bet struct {
	ID        string
	BankID    string
	Market    string
	Event     string
	EventDate time.Time
	Value     float64
	Result    float64
	Odd       float64
	CreatedAt time.Time
}

func NewBet(id, bankId, market, event string, eventDate time.Time, value, result, odd float64, createdAt time.Time) *Bet {
	return &Bet{
		ID:        id,
		BankID:    bankId,
		Market:    market,
		Event:     event,
		EventDate: eventDate,
		Value:     value,
		Result:    result,
		Odd:       odd,
		CreatedAt: createdAt,
	}
}

func (b *Bet) IsValidResult() bool {
	if b.Result >= 0 {
		return true
	}

	if b.Result+b.Value >= 0 {
		return true
	}

	return false
}
