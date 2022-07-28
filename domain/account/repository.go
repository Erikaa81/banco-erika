package account

import (
	"github.com/Erikaa81/banco-erika/domain/entities"
)

type Repository interface {
	CPFExists(string) bool
	Store(entities.Account) error
	List() ([]entities.Account, error)
	GetAccount(string) (entities.Account, error)
}
