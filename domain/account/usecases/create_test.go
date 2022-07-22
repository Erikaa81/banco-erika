package usecases

import (
	"errors"
	"testing"

	domainAccount "github.com/Erikaa81/banco-erika/domain/account"
	"github.com/Erikaa81/banco-erika/domain/entities"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		input := domainAccount.CreateAccountInput{
			Name:    "Carla",
			CPF:     "88877766555",
			PIN:     "34d",
			Balance: 0,
		}
		a := Account{
			repository: RepositoryMock{
				cpfExists: false,
				storeErr:  nil,
			},
		}

		got, err := a.CreateAccount(input)
		if err != nil {
			t.Errorf("wanted error to be nil but got: %s", err)
		}

		if got.ID == "" {
			t.Errorf("ID cannot be empty")
		}

		if got.Name != input.Name {
			t.Errorf("want %s, got,%s", got.Name, input.Name)
		}

		if got.CPF != input.CPF {
			t.Errorf("want %v, got,%v", input.CPF, got.CPF)
		}

		HashedPIN, _ := entities.HashPIN(input.PIN)
		if !entities.CheckPINHash(got.PIN, HashedPIN) {
			t.Errorf("pins do not match")
		}

		if got.Balance != input.Balance {
			t.Errorf("want, %v, got,%v", input.Balance, got.Balance)
		}

		if got.CreatedAt.IsZero() {
			t.Errorf("time cannot be zero")
		}
	})

	t.Run("should return error because the cpf already exists", func(t *testing.T) {
		c := domainAccount.CreateAccountInput{}
		a := Account{
			repository: RepositoryMock{
				cpfExists: true,
				storeErr:  nil,
			},
		}

		got, err := a.CreateAccount(c)
		want := entities.Account{}
		if got != want {
			t.Errorf("want empty account but got: %+v", got)
		}

		if !errors.Is(err, ErrAccountCreationDenied) {
			t.Errorf("want cpf already exists but got: %s", err)
		}
	})

	t.Run("should return store error", func(t *testing.T) {
		errStore := errors.New("something went wrong")
		c := domainAccount.CreateAccountInput{
			Name:    "Paulo",
			CPF:     "33344455566",
			PIN:     "fe4",
			Balance: 0,
		}
		a := Account{
			repository: RepositoryMock{
				cpfExists: false,
				storeErr:  errStore,
			},
		}
		got, err := a.CreateAccount(c)
		want := entities.Account{}
		if got != want {
			t.Errorf("want empty account but got: %+v", got)
		}

		if !errors.Is(err, errStore) {
			t.Errorf("expected store error and received: %s", err)
		}
	})
}

type RepositoryMock struct {
	cpfExists bool
	storeErr  error
}

func (r RepositoryMock) CPFExists(cpf string) bool {
	return r.cpfExists
}
func (r RepositoryMock) Store(account entities.Account) error {
	return r.storeErr
}
