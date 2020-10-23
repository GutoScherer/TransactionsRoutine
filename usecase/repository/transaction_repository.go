package repository

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"gorm.io/gorm"
)

// TransactionRepository represents a repository of transactions
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new TransactionRepository struct
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Save store the transaction in the database
func (repo TransactionRepository) Save(transaction *entity.Transaction) (*entity.Transaction, error) {
	err := repo.db.Create(transaction).Error
	return transaction, err
}
