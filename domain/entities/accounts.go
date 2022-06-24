package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Name      string
	Cpf       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

func CreateNewAccount(name, cpf, secret string, balance float64) (Account, error) {

	if name == "" {
		return Account{}, errors.New("é necessario um nome para criar a conta")
	}
	if cpf == "" {
		return Account{}, errors.New("é necessario um cpf para criar a conta")
	}
	if secret == "" {
		return Account{}, errors.New("uma senha é nessecario")
	}
	if balance < 0 {
		return Account{}, errors.New("o saldo da conta não pode ser negativo")
	}

	return Account{
		Id:        uuid.New().String(),
		Name:      name,
		Cpf:       cpf,
		Balance:   balance,
		Secret:    secret,
		CreatedAt: time.Now(),
	}, nil
}
