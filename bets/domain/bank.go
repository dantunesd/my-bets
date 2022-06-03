package domain

import (
	"time"
)

type IBanksRepository interface {
	CreateABank(Bank) error
	GetABank(id string) (Bank, error)
	UpdateABank(Bank) error
}

type Bank struct {
	ID           string    `bson:"_id" json:"id"`
	InitialValue float64   `json:"initial_value"`
	CurrentValue float64   `json:"current_value"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewBank(id string, initialValue float64, createdAt time.Time) *Bank {
	return &Bank{
		ID:           id,
		InitialValue: initialValue,
		CurrentValue: initialValue,
		CreatedAt:    createdAt,
	}
}

func (b *Bank) UpdateBank(value float64, updatedAt time.Time) {
	b.CurrentValue = b.CurrentValue + value
	b.UpdatedAt = updatedAt
}

func (b *Bank) isValidForFreeBets() bool {
	return b.CurrentValue >= 500
}
