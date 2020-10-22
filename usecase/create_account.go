package usecase

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
)

// CreateAccount represents a create account usecase and contains the dependencies to create an account
type CreateAccount struct {
	repo repository.AccountRepository
}

// NewCreateAccount creates a new CreateAccount struct
func NewCreateAccount(repo repository.AccountRepository) *CreateAccount {
	return &CreateAccount{repo: repo}
}

// Create creates and store an account on database
func (ca CreateAccount) Create(documentNumber string) (*entity.Account, error) {
	acc := entity.NewAccount(documentNumber)

	acc, err := ca.repo.Save(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
