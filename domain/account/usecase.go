package account

import (
	"github.com/Erikaa81/banco-erika/domain/entities"
)

type Usecase interface {
	CreateAccount(CreateAccountInput) (entities.Account, error)
	List() ([]entities.Account, error)
	GetAccount(string)(entities.Account, error)
}

type CreateAccountInput struct {
	Name    string
	CPF     string
	PIN     string
	Balance int
}

