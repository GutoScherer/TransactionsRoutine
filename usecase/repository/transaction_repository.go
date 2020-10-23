package repository

import (
	"fmt"

	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/go-sql-driver/mysql"
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
	if err != nil {
		fmt.Println(err)
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			return nil, buildRepositoryError(mysqlError)
		}
	}
	return transaction, nil
}
