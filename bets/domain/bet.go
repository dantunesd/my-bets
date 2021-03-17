package domain

import "time"

type Bet struct {
	ID        string    `bson:"_id" json:"id"`
	BankID    string    `json:"bank_id"`
	Market    string    `json:"market"`
	Event     string    `json:"event"`
	EventDate time.Time `json:"event_date"`
	Value     float64   `json:"value"`
	Result    float64   `json:"result"`
	Odd       float64   `json:"odd"`
	CreatedAt time.Time `json:"created_at"`
	free      bool
}

func NewBet(id, bankId, market, event string, eventDate time.Time, value, result, odd float64, createdAt time.Time, free bool) *Bet {
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
		free:      false,
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
