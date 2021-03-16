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
			b := &Bet{
				ID:        tt.fields.ID,
				BankID:    tt.fields.BankID,
				Market:    tt.fields.Market,
				Event:     tt.fields.Event,
				EventDate: tt.fields.EventDate,
				Value:     tt.fields.Value,
				Result:    tt.fields.Result,
				Odd:       tt.fields.Odd,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := b.IsValidResult(); got != tt.want {
				t.Errorf("Bet.IsValidResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
