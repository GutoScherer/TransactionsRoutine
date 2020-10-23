package repository

import (
	"errors"

	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/go-sql-driver/mysql"
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
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			return nil, buildRepositoryError(mysqlError)
		}
	}
	return acc, nil
}

// FindOneByID retrieves one account of the database by its ID
func (repo AccountRepository) FindOneByID(accountID uint64) (*entity.Account, error) {
	var acc entity.Account
	err := repo.db.First(&acc, accountID).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, ErrRegisterNotFound
	}

	return &acc, nil
}
