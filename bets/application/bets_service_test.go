package application

import (
	"errors"
	"my-bets/bets/domain/bank"
	"my-bets/bets/domain/bets"
	"testing"
)

func TestBetsService_PlaceABet(t *testing.T) {
	type fields struct {
		PbService      bets.IPlaceABetService
		BankRepository bank.IBanksRepository
		BetRepository  bets.IBetsRepository
	}
	type args struct {
		pbd PlaceBetDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bets.Bet
		wantErr bool
	}{
		{
			name: "should return an error if BankRepository.GetABank fails",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, errors.New("something went wrong")
					},
				},
			},
			args: args{
				pbd: PlaceBetDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return an error if PbService.PlaceABet fails",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
				},
				PbService: &PlaceABetServiceMock{
					PlaceABetMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
			},
			args: args{
				pbd: PlaceBetDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return an error if BetRepository.CreateABet fails",
			fields: fields{
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
				},
				PbService: &PlaceABetServiceMock{
					PlaceABetMockReturn: func() error {
						return nil
					},
				},
				BetRepository: &BetsRepositoryMock{
					CreateABetMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
			},
			args: args{
				pbd: PlaceBetDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return an error if BankRepository.UpdateABank fails",
			fields: fields{
				PbService: &PlaceABetServiceMock{
					PlaceABetMockReturn: func() error {
						return nil
					},
				},
				BetRepository: &BetsRepositoryMock{
					CreateABetMockReturn: func() error {
						return nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
					UpdateABankMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
			},
			args: args{
				pbd: PlaceBetDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a complete bet with an ID if not occurs any fails ",
			fields: fields{
				PbService: &PlaceABetServiceMock{
					PlaceABetMockReturn: func() error {
						return nil
					},
				},
				BetRepository: &BetsRepositoryMock{
					CreateABetMockReturn: func() error {
						return nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
					UpdateABankMockReturn: func() error {
						return nil
					},
				},
			},
			args: args{
				pbd: PlaceBetDTO{},
			},
			want:    &bets.Bet{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBetsService(
				tt.fields.PbService,
				tt.fields.BankRepository,
				tt.fields.BetRepository,
			)
			got, err := b.PlaceABet(tt.args.pbd)
			if (err != nil) != tt.wantErr {
				t.Errorf("BetsService.PlaceABet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if got == nil || got.ID == "" {
					t.Errorf("BetsService.PlaceABet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestBetsService_UndoABet(t *testing.T) {
	type fields struct {
		PbService      bets.IPlaceABetService
		BankRepository bank.IBanksRepository
		BetRepository  bets.IBetsRepository
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return error if BetRepository.GetABet fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, errors.New("something went wrong")
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: true,
		},
		{
			name: "should return error if BankRepository.GetABank fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, errors.New("something went wrong")
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: true,
		},
		{
			name: "should return error if PbService.UndoABet fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
				},
				PbService: &PlaceABetServiceMock{
					UndoABetMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: true,
		},
		{
			name: "should return error if BetRepository.DeleteABet fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, nil
					},
					DeleteABetMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
				},
				PbService: &PlaceABetServiceMock{
					UndoABetMockReturn: func() error {
						return nil
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: true,
		},
		{
			name: "should return error if BankRepository.UpdateABank fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, nil
					},
					DeleteABetMockReturn: func() error {
						return nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
					UpdateABankMockReturn: func() error {
						return errors.New("something went wrong")
					},
				},
				PbService: &PlaceABetServiceMock{
					UndoABetMockReturn: func() error {
						return nil
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: true,
		},
		{
			name: "should return nil if all functions not fails",
			fields: fields{
				BetRepository: &BetsRepositoryMock{
					GetABetMockReturn: func() (bets.Bet, error) {
						return bets.Bet{}, nil
					},
					DeleteABetMockReturn: func() error {
						return nil
					},
				},
				BankRepository: &BankRepositoryMock{
					GetABankMockReturn: func() (bank.Bank, error) {
						return bank.Bank{}, nil
					},
					UpdateABankMockReturn: func() error {
						return nil
					},
				},
				PbService: &PlaceABetServiceMock{
					UndoABetMockReturn: func() error {
						return nil
					},
				},
			},
			args: args{
				ID: "id",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BetsService{
				PbService:      tt.fields.PbService,
				BankRepository: tt.fields.BankRepository,
				BetRepository:  tt.fields.BetRepository,
			}
			if err := b.UndoABet(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("BetsService.UndoABet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type PlaceABetServiceMock struct {
	PlaceABetMockReturn func() error
	UndoABetMockReturn  func() error
}

func (p *PlaceABetServiceMock) PlaceABet(bet bets.Bet, bank *bank.Bank) error {
	return p.PlaceABetMockReturn()
}

func (p *PlaceABetServiceMock) UndoABet(bet bets.Bet, bank *bank.Bank) error {
	return p.UndoABetMockReturn()
}

type BetsRepositoryMock struct {
	CreateABetMockReturn func() error
	GetABetMockReturn    func() (bets.Bet, error)
	DeleteABetMockReturn func() error
}

func (b *BetsRepositoryMock) CreateABet(bets.Bet) error {
	return b.CreateABetMockReturn()
}

func (b *BetsRepositoryMock) GetABet(id string) (bets.Bet, error) {
	return b.GetABetMockReturn()
}

func (b *BetsRepositoryMock) DeleteABet(id string) error {
	return b.DeleteABetMockReturn()
}
