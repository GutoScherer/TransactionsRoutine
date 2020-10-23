package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//AccountRepository defines an interface for accounts persistence
type AccountRepository interface {
	Save(*entity.Account) (*entity.Account, error)
	FindOneByID(uint64) (*entity.Account, error)
}

type accountRepositoryMock struct{}

// NewAccountRepositoryMock creates a new AccountRepository implementation for unit tests
func NewAccountRepositoryMock() AccountRepository {
	return &accountRepositoryMock{}
}

// Save mocks the Save behavior
func (arm accountRepositoryMock) Save(acc *entity.Account) (*entity.Account, error) {
	return nil, nil
}

// FindOneByID mocks the FindOneByID behavior
func (arm accountRepositoryMock) FindOneByID(id uint64) (*entity.Account, error) {
	return nil, nil
}
