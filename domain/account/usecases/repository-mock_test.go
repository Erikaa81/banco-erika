package usecases

import "github.com/Erikaa81/banco-erika/domain/entities"

type RepositoryMock struct {
	cpfExists   bool
	storeErr    error
	accountList ([]entities.Account)
	listErr     error
	searchAccount (entities.Account)
	searchErr error
}

func (r RepositoryMock) CPFExists(cpf string) bool {
	return r.cpfExists
}
func (r RepositoryMock) Store(account entities.Account) error {
	return r.storeErr
}

func (r RepositoryMock) List() ([]entities.Account, error) {
	return r.accountList, r.listErr
}
func (r RepositoryMock) Get(id string) (entities.Account, error){
	return r.searchAccount, r.searchErr
}
