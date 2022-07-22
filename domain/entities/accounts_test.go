package entities

import (
	"reflect"
	"testing"
)

func TestNewAccount(t *testing.T) {
	t.Run("should return success when creating account", func(t *testing.T) {
		name := "Joana"
		cpf := "22244455544"
		pin := "12d"
		balance := 1000

		got, err := NewAccount(name, cpf, pin, balance)

		if got.ID == "" {
			t.Errorf("ID cannot be empty")
		}

		if got.Name != name {
			t.Errorf("want, %v, got,%v", name, got.Name)
		}

		if got.CPF != cpf {
			t.Errorf("want, %v, got,%v", cpf, got.CPF)
		}

		HashedPIN, _ := HashPIN(pin)
		if !CheckPINHash(got.PIN, HashedPIN) {
			t.Errorf("pins do not match")
		}

		if got.Balance != balance {
			t.Errorf("want, %v, got,%v", balance, got.Balance)
		}

		if got.CreatedAt.IsZero() {
			t.Errorf("time cannot be zero")
		}

		if err != nil {
			t.Errorf("wanted error to be nil but got: %s", err)
		}
	})

	type args struct {
		name    string
		cpf     string
		pin     string
		balance int
	}
	tests := []struct {
		name    string
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "should return error when creating account, the name is mandatory",
			args: args{
				name:    "",
				cpf:     "22244455544",
				pin:     "12d",
				balance: 1000,
			},
			want:    Account{},
			wantErr: true,
		},
		{
			name: "should return error when creating account, the CPF is mandatory",
			args: args{
				name:    "Paula",
				cpf:     " ",
				pin:     "45F",
				balance: -1000,
			},
			want:    Account{},
			wantErr: true,
		},
		{
			name: "should return error when creating account, the PIN is mandatory",
			args: args{
				name:    "Paula",
				cpf:     "55544477788",
				pin:     "",
				balance: 3000,
			},
			want:    Account{},
			wantErr: true,
		},
		{
			name: "should return err when creating account, balance cannot be negative",
			args: args{
				name:    "João",
				cpf:     "55233322211",
				pin:     "E45r",
				balance: -3,
			},
			want:    Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccount(tt.args.name, tt.args.cpf, tt.args.pin, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
