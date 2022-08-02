package usecases

import (
	"errors"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

var (
	ErrEmptyID = errors.New("the id was not filled")
)

func (a Account) Get(id string) (entities.Account, error) {
	if id == "" {
		return entities.Account{}, ErrEmptyID
	}

	account, err := a.repository.Get(id)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
