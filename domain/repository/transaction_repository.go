package repository

import "github.com/GutoScherer/TransactionsRoutine/domain/entity"

//TransactionRepository defines an interface for transaction persistence
type TransactionRepository interface {
	Save(*entity.Transaction) (*entity.Transaction, error)
}
