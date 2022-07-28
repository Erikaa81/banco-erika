package usecases

import (
	"errors"
	"testing"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

func TestAccount_GetAccount(t *testing.T) {
	t.Run("should get account id successfuly", func(t *testing.T) {
		account1, _ := entities.NewAccount("Paulo", "12343456785", "Er5", 10000)

		a := Account{
			repository: RepositoryMock{
				searchAccount: account1,
				searchErr:     nil,
			},
		}
		got, err := a.GetAccount(account1.ID)
		if got != account1 {
			t.Errorf("want: %v got: %v", account1, got)

			if err != nil {
				t.Errorf("wanted error to be nil and got: %v", err)
			}
		}
	})

	t.Run("error when looking for id", func(t *testing.T) {
		ErrSearch := errors.New("error when looking for account")
		want := entities.Account{}

		a := Account{
			repository: RepositoryMock{
				searchAccount: want,
				searchErr:     ErrSearch,
			},
		}
		got, err := a.GetAccount("")
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)

			if !errors.Is(err, ErrSearch) {
				t.Errorf("expected err%s, and received: %s", ErrSearch, err)
			}
		}
	})

	t.Run("should return error empty id", func(t *testing.T) {
		want := entities.Account{}

		a := Account{
			repository: RepositoryMock{
				searchAccount: want,
				searchErr:     nil,
			},
		}
		got, err := a.GetAccount("")
		if got != want {
			t.Errorf("want: %v, got :%v", want, got)
		}
		if !errors.Is(err, ErrEmptyId) {
			t.Errorf("expected err %s and got %s", ErrEmptyId, err)
		}
	})
}
