package bets

import (
	"my-bets/bets/domain/bank"
	"reflect"
	"testing"
	"time"
)

func TestPlaceABet(t *testing.T) {
	initialDate := time.Now()
	finalDate := initialDate.Add(time.Second * 10)

	type args struct {
		bet  Bet
		bank *bank.Bank
	}
	tests := []struct {
		name    string
		args    args
		want    *bank.Bank
		wantErr bool
	}{
		{
			name: "should update the bank",
			args: args{
				Bet{
					Result:    5,
					Value:     5,
					CreatedAt: finalDate,
				},
				&bank.Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 105,
				UpdatedAt:    finalDate,
			},
			wantErr: false,
		},
		{
			name: "should not update the bank when the bet is free and the bank is equal or higher than 500 and the result is negative",
			args: args{
				Bet{
					Result:    -5,
					Value:     5,
					CreatedAt: finalDate,
					Free:      true,
				},
				&bank.Bank{
					CurrentValue: 500,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 500,
				UpdatedAt:    initialDate,
			},
			wantErr: false,
		},
		{
			name: "should update the bank when the bet is free and the bank is higher than 500 and the result is positive",
			args: args{
				Bet{
					Result:    5,
					Value:     5,
					CreatedAt: finalDate,
					Free:      true,
				},
				&bank.Bank{
					CurrentValue: 500,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 505,
				UpdatedAt:    finalDate,
			},
			wantErr: false,
		},
		{
			name: "should return an error and not update the bank if the result is not valid",
			args: args{
				Bet{
					Value:     5,
					Result:    -10,
					CreatedAt: finalDate,
				},
				&bank.Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 100,
				UpdatedAt:    initialDate,
			},
			wantErr: true,
		},
		{
			name: "should return an error and not update the bank if the bet value is higher than the bank currentValue",
			args: args{
				Bet{
					Value:     150,
					Result:    0,
					CreatedAt: finalDate,
				},
				&bank.Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 100,
				UpdatedAt:    initialDate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaceABetService{}
			err := p.PlaceABet(tt.args.bet, tt.args.bank)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaceABet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.bank, tt.want) {
				t.Errorf("PlaceABet() = %v, want %v", tt.args.bank, tt.want)
			}
		})
	}
}

func TestPlaceABetService_UndoABet(t *testing.T) {
	initialDate := time.Now()
	finalDate := initialDate.Add(time.Second * 10)

	type args struct {
		bet  Bet
		bank *bank.Bank
	}
	tests := []struct {
		name    string
		p       *PlaceABetService
		args    args
		want    *bank.Bank
		wantErr bool
	}{
		{
			name: "should update to bank undoing the result",
			args: args{
				Bet{
					Result:    5,
					Value:     5,
					CreatedAt: finalDate,
				},
				&bank.Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: &bank.Bank{
				CurrentValue: 95,
				UpdatedAt:    finalDate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaceABetService{}
			if err := p.UndoABet(tt.args.bet, tt.args.bank); (err != nil) != tt.wantErr {
				t.Errorf("PlaceABetService.UndoABet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.bank, tt.want) {
				t.Errorf("PlaceABet() = %v, want %v", tt.args.bank, tt.want)
			}
		})
	}
}
