package usecases

import "github.com/Erikaa81/banco-erika/domain/account"

type Account struct {
	repository account.Repository
}

func NewAccount(repository account.Repository) Account {
	return Account{repository: repository}
}
