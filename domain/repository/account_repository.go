package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//AccountRepository defines an interface for accounts persistence
type AccountRepository interface {
	Save(*entity.Account) (*entity.Account, error)
	FindOneByID(uint64) (*entity.Account, error)
}

type accountRepositoryMock struct {
	acc *entity.Account
	err error
}

// NewAccountRepositoryMock creates a new AccountRepository implementation for unit tests
func NewAccountRepositoryMock(acc *entity.Account, err error) AccountRepository {
	return &accountRepositoryMock{acc: acc, err: err}
}

// Save mocks the Save behavior
func (arm accountRepositoryMock) Save(acc *entity.Account) (*entity.Account, error) {
	if arm.err != nil {
		return nil, arm.err
	}

	return arm.acc, nil
}

// FindOneByID mocks the FindOneByID behavior
func (arm accountRepositoryMock) FindOneByID(id uint64) (*entity.Account, error) {
	if arm.err != nil {
		return nil, arm.err
	}

	return arm.acc, nil
}
