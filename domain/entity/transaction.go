package entity

import (
	"fmt"
	"time"
)

// Transaction is an entity for transactions data storage
type Transaction struct {
	ID              uint64
	AccountID       uint64
	Account         Account
	OperationTypeID OperationType
	Amount          float64
	CreatedAt       time.Time
}

// NewTransaction creates a new Transaction struct
func NewTransaction(accountID uint64, operationTypeID uint64, amount float64) (*Transaction, error) {
	operationType, err := NewOperationType(operationTypeID)
	if err != nil {
		return nil, fmt.Errorf("NewTransaction error: %v", err)
	}

	account := Account{ID: accountID}

	if operationType.IsDebit() {
		amount = -amount
	}

	return &Transaction{
		Account:         account,
		OperationTypeID: operationType,
		Amount:          amount,
	}, nil
}
