package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//AccountRepository defines an interface for accounts persistence
type AccountRepository interface {
	Save(*entity.Account) (*entity.Account, error)
	FindOneByID(uint64) (*entity.Account, error)
}
