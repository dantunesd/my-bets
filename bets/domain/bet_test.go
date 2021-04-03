package domain

import (
	"testing"
	"time"
)

func TestBet_IsValidResult(t *testing.T) {
	type fields struct {
		ID        string
		BankID    string
		Market    string
		Event     string
		EventDate time.Time
		Value     float64
		Result    float64
		Odd       float64
		free      bool
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return true when the Bet Value is Higher than Result",
			fields: fields{
				Value:  5,
				Result: 0,
			},
			want: true,
		},
		{
			name: "should return true when the Bet Value is Lower than Result",
			fields: fields{
				Value:  5,
				Result: 10,
			},
			want: true,
		},
		{
			name: "should return true when the Bet Result is the Bet Value but negative",
			fields: fields{
				Value:  5,
				Result: -5,
			},
			want: true,
		},
		{
			name: "should return true when the Bet Result is negative",
			fields: fields{
				Value:  5,
				Result: -1,
			},
			want: true,
		},
		{
			name: "should return false when the Bet Result is more negative than it can be",
			fields: fields{
				Value:  5,
				Result: -6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBet(
				tt.fields.ID,
				tt.fields.BankID,
				tt.fields.Market,
				tt.fields.Event,
				tt.fields.EventDate,
				tt.fields.Value,
				tt.fields.Result,
				tt.fields.Odd,
				tt.fields.free,
				tt.fields.CreatedAt,
			)
			if got := b.IsValidResult(); got != tt.want {
				t.Errorf("Bet.IsValidResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
