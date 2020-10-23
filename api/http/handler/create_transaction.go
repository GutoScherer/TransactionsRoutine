package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
)

type createTransactionRequest struct {
	AccountID       uint64  `json:"account_id"`
	OperationTypeID uint64  `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

// TransactionCreator represents an creator for transaction usecases
type TransactionCreator interface {
	Create(accountID, operationTypeID uint64, amount float64) (*presenter.CreateTransactionOutput, error)
}

// CreateTransaction represents the http handler struct to create an transaction
type CreateTransaction struct {
	transactionCreator TransactionCreator
}

// NewCreateTransaction creates a new CreateTransaction struct
func NewCreateTransaction(tc TransactionCreator) *CreateTransaction {
	return &CreateTransaction{transactionCreator: tc}
}

// HandlerFunc is the http handler function used by the server
func (h CreateTransaction) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	var createTransactionRequest createTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
	if err != nil {
		// TODO: Error handling
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.transactionCreator.Create(createTransactionRequest.AccountID, createTransactionRequest.OperationTypeID, createTransactionRequest.Amount)
	if err != nil {
		// TODO: Error handling
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(output)
}
