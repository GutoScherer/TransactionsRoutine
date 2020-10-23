package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
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
	logger             *log.Logger
}

// NewCreateTransaction creates a new CreateTransaction struct
func NewCreateTransaction(tc TransactionCreator, logger *log.Logger) *CreateTransaction {
	return &CreateTransaction{
		transactionCreator: tc,
		logger:             logger,
	}
}

// HandlerFunc is the http handler function used by the server
func (h CreateTransaction) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	var createTransactionRequest createTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
	if err != nil {
		h.logger.Println("invalid request body:", err)

		output := map[string]string{"error": "invalid JSON body"}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(output)
		return
	}

	output, err := h.transactionCreator.Create(createTransactionRequest.AccountID, createTransactionRequest.OperationTypeID, createTransactionRequest.Amount)
	if err != nil {
		h.logger.Println("error creating transaction:", err)

		if errors.Is(err, repository.ErrInvalidData) || errors.Is(err, repository.ErrForeignKeyConstraint) {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(output)
}
