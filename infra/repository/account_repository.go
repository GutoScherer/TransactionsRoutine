package repository

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"gorm.io/gorm"
)

// AccountRepository represents a repository of accounts
type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository creates a new AccountRepository struct
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Save store the account in the database
func (repo AccountRepository) Save(acc *entity.Account) (*entity.Account, error) {
	err := repo.db.Create(acc).Error
	return acc, err
}

// FindOneByID retrieves one account of the database by its ID
func (repo AccountRepository) FindOneByID(accountID uint64) (*entity.Account, error) {
	var acc entity.Account
	err := repo.db.First(&acc, accountID).Error
	return &acc, err
}
