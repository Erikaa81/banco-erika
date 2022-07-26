package usecases

import (
	"errors"

	"github.com/Erikaa81/banco-erika/domain/account"
	"github.com/Erikaa81/banco-erika/domain/entities"
)

var (
	ErrCPFAlreadyExists = errors.New("there is already an account for this cpf")
)

func (a Account) CreateAccount(input account.CreateAccountInput) (entities.Account, error) {
	if a.repository.CPFExists(input.CPF) {
		return entities.Account{}, ErrCPFAlreadyExists
	}

	account, err := entities.NewAccount(input.Name, input.CPF, input.PIN, input.Balance)
	if err != nil {
		return entities.Account{}, err
	}

	err = a.repository.Store(account)
	if err != nil {
		return entities.Account{}, err

	}
	return account, nil
}
