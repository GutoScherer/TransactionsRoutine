package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//TransactionRepository defines an interface for transaction persistence
type TransactionRepository interface {
	Save(*entity.Transaction) (*entity.Transaction, error)
}

type transactionRepositoryMock struct{}

// NewTransactionRepositoryMock creates a new TransactionRepository implementation for unit tests
func NewTransactionRepositoryMock() TransactionRepository {
	return &transactionRepositoryMock{}
}

// Save mocks the Save behavior
func (trm transactionRepositoryMock) Save(_ *entity.Transaction) (*entity.Transaction, error) {
	return nil, nil
}
