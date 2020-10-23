package presenter

import (
	"time"

	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
)

// TransactionPresenter represents the output port for transaction usecases
type TransactionPresenter interface {
	CreateTransactionOutput(*entity.Transaction) *CreateTransactionOutput
}

type transactionPresenter struct{}

// NewTransactionPresenter creates a new TransactionPresenter implementation
func NewTransactionPresenter() TransactionPresenter {
	return &transactionPresenter{}
}

func (tp transactionPresenter) CreateTransactionOutput(transaction *entity.Transaction) *CreateTransactionOutput {
	return &CreateTransactionOutput{
		AccountID:     transaction.AccountID,
		OperationType: transaction.OperationType.String(),
		Amount:        transaction.Amount,
		CreatedAt:     transaction.CreatedAt,
	}
}

// CreateTransactionOutput represents the output data for the create transaction usecase
type CreateTransactionOutput struct {
	AccountID     uint64    `json:"account_id"`
	OperationType string    `json:"operation_type"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
