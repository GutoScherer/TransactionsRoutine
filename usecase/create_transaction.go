package usecase

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
)

// CreateTransaction represents a create transaction usecase and contains the dependencies to create a transaction
type CreateTransaction struct {
	repo repository.TransactionRepository
}

// NewCreateTransaction creates a new CreateTransaction struct
func NewCreateTransaction(repo repository.TransactionRepository) *CreateTransaction {
	return &CreateTransaction{repo: repo}
}

// Create creates and store a transaction on database
func (ct CreateTransaction) Create(accountID, operationTypeID uint64, amount float64) (*entity.Transaction, error) {
	transaction, err := entity.NewTransaction(accountID, operationTypeID, amount)
	if err != nil {
		return nil, err
	}

	transaction, err = ct.repo.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
