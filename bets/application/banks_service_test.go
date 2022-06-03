package application

import (
	"errors"
	"my-bets/bets/domain"
	"reflect"
	"testing"
	"time"
)

func TestBanksService_CreateABank(t *testing.T) {
	bank := domain.Bank{
		ID:           "id",
		InitialValue: 200,
		CurrentValue: 200,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	type fields struct {
		BankRepository domain.IBanksRepository
	}
	type args struct {
		initialValue float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Bank
		wantErr bool
	}{
		{
			name: "should create a bank and return it",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					CreateABankMockReturn: func() error {
						return nil
					},
				},
			},
			args: args{
				initialValue: 200,
			},
			want:    bank,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBankService(
				tt.fields.BankRepository,
			)

			got, err := b.CreateABank(tt.args.initialValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("BanksService.CreateABank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.InitialValue != tt.want.InitialValue {
				t.Errorf("BanksService.CreateABank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanksService_GetABank(t *testing.T) {
	bank := domain.Bank{
		ID:           "id",
		InitialValue: 200,
		CurrentValue: 200,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	type fields struct {
		BankRepository domain.IBanksRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Bank
		wantErr bool
	}{
		{
			name: "should returns the bank retrieved from repository",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (domain.Bank, error) {
						return bank, nil
					},
				},
			},
			args: args{
				id: "id",
			},
			want:    bank,
			wantErr: false,
		},
		{
			name: "should returns the error the retrieved from repository",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (domain.Bank, error) {
						return domain.Bank{}, errors.New("failed to get data from db")
					},
				},
			},
			args: args{
				id: "id",
			},
			want:    domain.Bank{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BanksService{
				BankRepository: tt.fields.BankRepository,
			}
			got, err := b.GetABank(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BanksService.GetABank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BanksService.GetABank() = %v, want %v", got, tt.want)
			}
		})
	}
}

type BankRepositoryMock struct {
	CreateABankMockReturn func() error
	GetABankMockReturn    func() (domain.Bank, error)
	UpdateABankMockReturn func() error
}

func (b *BankRepositoryMock) CreateABank(bank domain.Bank) error {
	return b.CreateABankMockReturn()
}

func (b *BankRepositoryMock) GetABank(id string) (domain.Bank, error) {
	return b.GetABankMockReturn()
}

func (b *BankRepositoryMock) UpdateABank(bank domain.Bank) error {
	return b.UpdateABankMockReturn()
}
