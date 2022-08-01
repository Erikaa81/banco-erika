package usecases

import (
	"errors"
	"testing"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

func TestAccount_GetAccount(t *testing.T) {
	t.Run("should successfully get account by id", func(t *testing.T) {
		expectedAccount, _ := entities.NewAccount("Paulo", "12343456785", "Er5", 10000)

		a := Account{
			repository: RepositoryMock{
				searchAccount: expectedAccount,
				searchErr:     nil,
			},
		}
		got, err := a.Get(expectedAccount.ID)
		if got != expectedAccount {
			t.Errorf("want: %v got: %v", expectedAccount, got)

			if err != nil {
				t.Errorf("wanted error to be nil and got: %v", err)
			}
		}
	})

	t.Run("should return error when looking for id at repository", func(t *testing.T) {
		ErrSearch := errors.New("error when looking for account")
		want := entities.Account{}

		a := Account{
			repository: RepositoryMock{
				searchAccount: want,
				searchErr:     ErrSearch,
			},
		}
		got, err := a.Get("")
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)

			if !errors.Is(err, ErrSearch) {
				t.Errorf("expected err%s, and received: %s", ErrSearch, err)
			}
		}
	})

	t.Run("should return error when empty id is informed", func(t *testing.T) {
		want := entities.Account{}

		a := Account{
			repository: RepositoryMock{
				searchAccount: want,
				searchErr:     nil,
			},
		}
		got, err := a.Get("")
		if got != want {
			t.Errorf("want: %v, got :%v", want, got)
		}
		if !errors.Is(err, ErrEmptyID) {
			t.Errorf("expected err %s and got %s", ErrEmptyID, err)
		}
	})
}
