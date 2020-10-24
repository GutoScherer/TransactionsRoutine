package entity

import (
	"fmt"

	"github.com/GutoScherer/TransactionsRoutine/domain"
)

// OperationType is an enum that represents the operation types of transactions
//
// Valid values are:
//  1 - Compra à Vista
//  2 - Compra Parcelada
//  3 - Saque
//  4 - Pagamento
type OperationType uint64

const (
	//CompraAVista representa uma compra a vista
	CompraAVista OperationType = 1 + iota
	//CompraParcelada representa uma compra parcelada
	CompraParcelada
	//Saque representa um saque
	Saque
	//Pagamento representa um pagamento
	Pagamento
	end
)

var types = [...]string{
	"COMPRA A VISTA",
	"COMPRA PARCELADA",
	"SAQUE",
	"PAGAMENTO",
}

// NewOperationType creates a new operation type
func NewOperationType(operationTypeID uint64) (OperationType, error) {
	ot := OperationType(operationTypeID)
	if !ot.IsValid() {
		return 0, domain.NewDomainError(fmt.Sprintf("invalid operation type '%d'", operationTypeID))
	}

	return ot, nil
}

func (ot OperationType) String() string {
	return types[ot-1]
}

//IsValid checks if the operation type is valid
func (ot OperationType) IsValid() bool {
	return ot < end
}

//IsDebit checks if the operation type is a debit operation
func (ot OperationType) IsDebit() bool {
	switch ot {
	case Pagamento:
		return false
	}

	return true
}
