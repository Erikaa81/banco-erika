package usecases

import "github.com/Erikaa81/banco-erika/domain/entities"

func (a Account) List() ([]entities.Account, error) {
	accounts, err := a.repository.List()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
