package usecase

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
)

// RetrieveAccount represents a retrieve account usecase and contains the dependencies to find an account
type RetrieveAccount struct {
	repo repository.AccountRepository
}

// NewRetrieveAccount creates a new RetrieveAccount struct
func NewRetrieveAccount(repo repository.AccountRepository) *RetrieveAccount {
	return &RetrieveAccount{repo: repo}
}

// RetrieveByID finds one account by its primary key
func (ra RetrieveAccount) RetrieveByID(accountID uint64) (*entity.Account, error) {
	acc, err := ra.repo.FindOneByID(accountID)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
