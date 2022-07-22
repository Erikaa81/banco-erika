package account

import (
	"github.com/Erikaa81/banco-erika/domain/entities"
)

type AccountUsecase interface {
	CreateAccount(CreateAccountInput) (entities.Account, error)
}

type CreateAccountInput struct {
	Name    string
	CPF     string
	PIN     string
	Balance int
}
