package entities

import (
	"reflect"
	"testing"
)

func TestNewTransfer(t *testing.T) {
	t.Run("should return sucess transfer", func(t *testing.T) {
		accountOriginID := "4555"
		accountDestinationID := "26784"
		amount := 2000

		got, err := NewTransfer(accountOriginID, accountDestinationID, amount)

		if got.ID == "" {
			t.Errorf("ID cannot to be zero")
		}

		if got.AccountOriginID != accountOriginID {
			t.Errorf("want, %v, got,%v", accountOriginID, got)
		}

		if got.Amount != amount {
			t.Errorf("want, %v, got,%v", amount, got)
		}

		if got.CreatedAt.IsZero() {
			t.Errorf("createdat cannot to be zero")
		}

		if err != nil {
			t.Errorf("wanted error to be nil but got: %s", err)
		}
	})

	type args struct {
		accountOriginID      string
		accountDestinationID string
		amount               int
	}
	tests := []struct {
		name    string
		args    args
		want    Transfer
		wantErr bool
	}{
		{
			name: "should return transfer error because the account origin is mandatory",
			args: args{
				accountOriginID:      "",
				accountDestinationID: "45667",
				amount:               10,
			},
			want:    Transfer{},
			wantErr: true,
		},
		{
			name: "should return transfer error because the account destination is mandatory",
			args: args{
				accountOriginID:      "65789",
				accountDestinationID: "",
				amount:               10,
			},
			want:    Transfer{},
			wantErr: true,
		},
		{
			name: "should return transfer error because the accounts are the same",
			args: args{
				accountOriginID:      "123457",
				accountDestinationID: "123457",
				amount:               10,
			},
			want:    Transfer{},
			wantErr: true,
		},
		{
			name: "shold return err invalid transfer as the amount to be transferred is not valid",
			args: args{
				accountOriginID:      "4333",
				accountDestinationID: "4321",
				amount:               -10,
			},
			want:    Transfer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransfer(tt.args.accountOriginID, tt.args.accountDestinationID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
