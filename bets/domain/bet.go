package domain

import "time"

type IBetsRepository interface {
	CreateABet(Bet) error
	GetABet(id string) (Bet, error)
	DeleteABet(id string) error
}

type Bet struct {
	ID        string    `bson:"_id" json:"id"`
	BankID    string    `json:"bank_id"`
	Market    string    `json:"market"`
	Event     string    `json:"event"`
	EventDate time.Time `json:"event_date"`
	Value     float64   `json:"value"`
	Result    float64   `json:"result"`
	Odd       float64   `json:"odd"`
	Free      bool      `json:"free"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBet(id, bankId, market, event string, eventDate time.Time, value, result, odd float64, free bool, createdAt time.Time) *Bet {
	return &Bet{
		ID:        id,
		BankID:    bankId,
		Market:    market,
		Event:     event,
		EventDate: eventDate,
		Value:     value,
		Result:    result,
		Odd:       odd,
		Free:      free,
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

func (b *Bet) needApplyFreeBet() bool {
	return b.Free && b.Result < 0
}
