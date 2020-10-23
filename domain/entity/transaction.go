package entity

import (
	"fmt"
	"time"

	"github.com/GutoScherer/TransactionsRoutine/domain"
)

// Transaction is an entity for transactions data storage
type Transaction struct {
	ID            uint64
	AccountID     uint64
	Account       Account
	OperationType OperationType `gorm:"column:operation_type_id"`
	Amount        float64
	CreatedAt     time.Time
}

// NewTransaction creates a new Transaction struct
func NewTransaction(accountID uint64, operationTypeID uint64, amount float64) (*Transaction, error) {
	operationType, err := NewOperationType(operationTypeID)
	if err != nil {
		return nil, domain.NewDomainError(fmt.Sprintf("new transaction error: %v", err))
	}

	account := Account{ID: accountID}

	if operationType.IsDebit() {
		amount = -amount
	}

	return &Transaction{
		Account:       account,
		OperationType: operationType,
		Amount:        amount,
	}, nil
}
