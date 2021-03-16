package domain

import (
	"time"
)

type Bank struct {
	ID           string `bson:"_id" json:"id,omitempty"`
	InitialValue float64
	CurrentValue float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewBank(id string, initialValue, currentValue float64, createdAt, updatedAt time.Time) *Bank {
	return &Bank{
		ID:           id,
		InitialValue: initialValue,
		CurrentValue: currentValue,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
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
