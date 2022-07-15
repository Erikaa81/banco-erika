package entities

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewLogin(t *testing.T) {
	t.Run("should return login success", func(t *testing.T) {
		cpf := "22244465544"
		secret := "67a"

		got, err := NewLogin(cpf, secret)

		if got.CPF != cpf {
			t.Errorf("want, %v, got,%v", cpf, got.CPF)
		}

		HashedSecret, _ := HashSecret(secret)
		if CheckSecretHash(got.Secret, HashedSecret) != true {
			t.Errorf("wnat, %v, got, %v", HashedSecret, got.Secret)
		}

		if !errors.Is(err, nil) {
			t.Errorf("want, %v, got,%v ", nil, err)
		}
	})

	type args struct {
		cpf    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    Login
		wantErr bool
	}{
		{
			name: "should return error when logging in, because the cpf is mandatory",
			args: args{
				cpf:    "",
				secret: "uu877",
			},
			want:    Login{},
			wantErr: true,
		},
		{
			name: "should return error when logging in, because the secret is mandatory",
			args: args{
				cpf:    "88877744433",
				secret: "",
			},
			want:    Login{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLogin(tt.args.cpf, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
