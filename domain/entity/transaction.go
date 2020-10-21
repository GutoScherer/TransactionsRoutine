package entity

import (
	"fmt"
	"time"
)

//Transaction is an entity for transactions data storage
type Transaction struct {
	ID            uint64
	Account       Account
	OperationType OperationType
	Amount        float64
	CreatedAt     time.Time
}

//NewTransaction creates a new instance of Transaction
func NewTransaction(accountID uint64, operationTypeID uint64, amount float64) (*Transaction, error) {
	operationType := OperationType(operationTypeID)
	if !operationType.IsValid() {
		return nil, fmt.Errorf("Invalid operation type '%d'", operationTypeID)
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
