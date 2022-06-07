package bank

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBank(t *testing.T) {
	createdAt := time.Now()
	ID := "string-id"

	type args struct {
		ID           string
		initialValue float64
		createdAt    time.Time
	}
	tests := []struct {
		name string
		args args
		want *Bank
	}{
		{
			name: "should return a new Bank",
			args: args{ID, 100, createdAt},
			want: &Bank{
				ID:           ID,
				InitialValue: 100,
				CurrentValue: 100,
				CreatedAt:    createdAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBank(tt.args.ID, tt.args.initialValue, tt.args.createdAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_UpdateBank(t *testing.T) {
	initialDate := time.Now()
	finalDate := initialDate.Add(time.Second * 10)

	type fields struct {
		ID           string
		initialValue float64
		currentValue float64
		createdAt    time.Time
		updatedAt    time.Time
	}
	type args struct {
		value     float64
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Bank
	}{
		{
			name: "should sum the value inputed and update the time",
			fields: fields{
				currentValue: 100,
				updatedAt:    initialDate,
			},
			args: args{
				value:     50,
				updatedAt: finalDate,
			},
			want: &Bank{
				CurrentValue: 150,
				UpdatedAt:    finalDate,
			},
		},
		{
			name: "should subtract the value inputed and update the time",
			fields: fields{
				currentValue: 100,
				updatedAt:    initialDate,
			},
			args: args{
				value:     -50,
				updatedAt: finalDate,
			},
			want: &Bank{
				CurrentValue: 50,
				UpdatedAt:    finalDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bank{
				tt.fields.ID,
				tt.fields.initialValue,
				tt.fields.currentValue,
				tt.fields.createdAt,
				tt.fields.updatedAt,
			}
			if b.UpdateBank(tt.args.value, tt.args.updatedAt); !reflect.DeepEqual(b, tt.want) {
				t.Errorf("Bank.UpdateBank() = %v, want %v", b, tt.want)
			}
		})
	}
}
