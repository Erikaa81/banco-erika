package account

import (
	"github.com/Erikaa81/banco-erika/domain/entities"
)

type Repository interface {
	CPFExists(string) bool
	Store(entities.Account) error
}
