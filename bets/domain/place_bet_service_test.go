package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestPlaceABet(t *testing.T) {
	initialDate := time.Now()
	finalDate := initialDate.Add(time.Second * 10)

	type args struct {
		bet  Bet
		bank Bank
	}
	tests := []struct {
		name    string
		args    args
		want    Bank
		wantErr bool
	}{
		{
			name: "should return a bank updated",
			args: args{
				Bet{
					Result:    5,
					Value:     5,
					CreatedAt: finalDate,
				},
				Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: Bank{
				CurrentValue: 105,
				UpdatedAt:    finalDate,
			},
			wantErr: false,
		},
		{
			name: "should return an error and the bank not updated if the result is not valid",
			args: args{
				Bet{
					Value:     5,
					Result:    -10,
					CreatedAt: finalDate,
				},
				Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: Bank{
				CurrentValue: 100,
				UpdatedAt:    initialDate,
			},
			wantErr: true,
		},
		{
			name: "should return an error and the bank not updated if the bet value is higher than the bank currentValue",
			args: args{
				Bet{
					Value:     150,
					Result:    0,
					CreatedAt: finalDate,
				},
				Bank{
					CurrentValue: 100,
					UpdatedAt:    initialDate,
				},
			},
			want: Bank{
				CurrentValue: 100,
				UpdatedAt:    initialDate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaceABetService{}
			got, err := p.PlaceABet(tt.args.bet, tt.args.bank)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaceABet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaceABet() = %v, want %v", got, tt.want)
			}
		})
	}
}
