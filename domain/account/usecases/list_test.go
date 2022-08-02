package usecases

import (
	"errors"
	"testing"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

func TestList(t *testing.T) {
	t.Run("should successfully return account list", func(t *testing.T) {
		account1, err1 := entities.NewAccount("erika", "22233344455", "Er5", 1000)
		account2, err2 := entities.NewAccount("Maria", "55533344455", "F35", 2000)
		account3, err3 := entities.NewAccount("Paula", "44455566647", "E43", 3000)

		if err1 != nil || err2 != nil || err3 != nil {
			t.Fatalf("fix before running tests: all errors should be nil but got %s, %s, %s",err1, err2, err3)
		}

		want := []entities.Account{account1, account2, account3}
		a := Account{
			repository: RepositoryMock{
				accountList: want,
				listErr:     nil,
			},
		}

		got, err := a.List()
		for i, v := range got {
			if want[i] != v {
				t.Errorf("want:%+v got%+v", want[i], got[i])
			}
		}

		if err != nil {
			t.Errorf("wanted error to be nil and got: %v", err)
		}
	})

	t.Run("should successfully return an empty list", func(t *testing.T) {
		a := Account{
			repository: RepositoryMock{
				accountList: []entities.Account{},
				listErr:     nil,
			},
		}
		got, err := a.List()

		if len(got) != 0 {
			t.Errorf("wanted empty list but got: %+v", got)
		}

		if err != nil {
			t.Errorf("wanted error to be nil and got: %v", err)
		}

	})
	t.Run("should return error when listing", func(t *testing.T) {
		ErrList := errors.New("error when listing")

		a := Account{
			repository: RepositoryMock{
				accountList: nil,
				listErr:     ErrList,
			},
		}
		got, err := a.List()
		if len(got) != 0 {
			t.Errorf("wanted error while listing but got: %+v", got)
		}

		if !errors.Is(err, ErrList) {
			t.Errorf("expected err and received: %s", err)
		}
	})

}
