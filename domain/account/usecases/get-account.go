package usecases

import (
	"errors"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

var (
	ErrEmptyId = errors.New("the id was not filled")
)

func (a Account) GetAccount(id string) (entities.Account, error) {
	if id == "" {
		return entities.Account{}, ErrEmptyId
	}

	account, err := a.repository.GetAccount(id)
	if err != nil {
		return account, err
	}
	return account, nil
}
