package domain

import (
	"time"
)

type Bank struct {
	ID           string
	InitialValue float64
	CurrentValue float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func InitializeANewBank(id string, initialValue float64, createdAt time.Time) *Bank {
	return &Bank{
		ID:           id,
		InitialValue: initialValue,
		CurrentValue: initialValue,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}
}

func (b *Bank) UpdateBank(value float64, updatedAt time.Time) {
	b.CurrentValue = b.CurrentValue + value
	b.UpdatedAt = updatedAt
}
