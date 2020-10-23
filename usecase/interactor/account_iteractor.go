package interactor

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
)

// AccountInteractor represents an interactor for account usecases
type AccountInteractor interface {
	RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error)
	Create(documentNumber string) (*presenter.CreateAccountOutput, error)
}

type accountInteractor struct {
	AccountRepository repository.AccountRepository
	AccountPresenter  presenter.AccountPresenter
}

type accountInteractorMock struct{}

// NewAccountInteractor creates a new AccountInteractor implementation
func NewAccountInteractor(repo repository.AccountRepository, ap presenter.AccountPresenter) AccountInteractor {
	return &accountInteractor{AccountRepository: repo, AccountPresenter: ap}
}

// NewAccountInteractorMock creates a new AccountInteractor implementation for unit tests
func NewAccountInteractorMock() AccountInteractor {
	return &accountInteractorMock{}
}

// RetrieveByID finds one account by its primary key
func (ai accountInteractor) RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error) {
	acc, err := ai.AccountRepository.FindOneByID(accountID)
	if err != nil {
		return nil, err
	}

	return ai.AccountPresenter.RetrieveAccountOutput(acc), nil
}

// Create creates and store an account on database
func (ai accountInteractor) Create(documentNumber string) (*presenter.CreateAccountOutput, error) {
	acc := entity.NewAccount(documentNumber)

	acc, err := ai.AccountRepository.Save(acc)
	if err != nil {
		return nil, err
	}

	return ai.AccountPresenter.CreateAccountOutput(acc), nil
}

// RetrieveByID mocks the RetrieveByID behavior
func (aim accountInteractorMock) RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error) {
	return nil, nil
}

// Create mocks the Create behavior
func (aim accountInteractorMock) Create(documentNumber string) (*presenter.CreateAccountOutput, error) {
	return nil, nil
}
