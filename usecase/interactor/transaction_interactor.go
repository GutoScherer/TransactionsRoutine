package interactor

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
)

// TransactionInteractor represents an interactor for transaction usecases
type TransactionInteractor interface {
	Create(accountID, operationTypeID uint64, amount float64) (*presenter.CreateTransactionOutput, error)
}

type transactionInteractor struct {
	TransactionRepository repository.TransactionRepository
	TransactionPresenter  presenter.TransactionPresenter
}

// NewTransactionInteractor creates a new TransactionInteractor implementation
func NewTransactionInteractor(repo repository.TransactionRepository, tp presenter.TransactionPresenter) TransactionInteractor {
	return &transactionInteractor{TransactionRepository: repo, TransactionPresenter: tp}
}

// Create creates and store a transaction on database
func (ti transactionInteractor) Create(accountID, operationTypeID uint64, amount float64) (*presenter.CreateTransactionOutput, error) {
	transaction, err := entity.NewTransaction(accountID, operationTypeID, amount)
	if err != nil {
		return nil, err
	}

	transaction, err = ti.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return ti.TransactionPresenter.CreateTransactionOutput(transaction), nil
}
