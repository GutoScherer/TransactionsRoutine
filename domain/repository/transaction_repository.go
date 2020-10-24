package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//TransactionRepository defines an interface for transaction persistence
type TransactionRepository interface {
	Save(*entity.Transaction) (*entity.Transaction, error)
}

type transactionRepositoryMock struct {
	transaction *entity.Transaction
	err         error
}

// NewTransactionRepositoryMock creates a new TransactionRepository implementation for unit tests
func NewTransactionRepositoryMock(t *entity.Transaction, err error) TransactionRepository {
	return &transactionRepositoryMock{transaction: t, err: err}
}

// Save mocks the Save behavior
func (trm transactionRepositoryMock) Save(_ *entity.Transaction) (*entity.Transaction, error) {
	if trm.err != nil {
		return nil, trm.err
	}

	return trm.transaction, nil
}
