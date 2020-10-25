package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/domain"
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
	responseWriter := newResponseWriter(rw)

	err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
	if err != nil {
		h.logger.Println("invalid request body:", err)

		output := map[string]string{"error": "invalid JSON body"}
		responseWriter.outputResponse(http.StatusBadRequest, output)
		return
	}

	output, err := h.transactionCreator.Create(createTransactionRequest.AccountID, createTransactionRequest.OperationTypeID, createTransactionRequest.Amount)
	if err != nil {
		h.logger.Println("error creating transaction:", err)

		if _, ok := err.(*domain.Error); ok {
			responseWriter.outputResponse(http.StatusUnprocessableEntity, nil)
			return
		}

		if errors.Is(err, repository.ErrInvalidData) || errors.Is(err, repository.ErrForeignKeyConstraint) {
			responseWriter.outputResponse(http.StatusUnprocessableEntity, nil)
			return
		}

		responseWriter.outputResponse(http.StatusInternalServerError, nil)
	}

	responseWriter.outputResponse(http.StatusCreated, output)
	return
}

type transactionCreatorMock struct {
	output *presenter.CreateTransactionOutput
	err    error
}

// NewTransactionCreatorMock creates a new TransactionCreator implementation for unit tests
func NewTransactionCreatorMock(output *presenter.CreateTransactionOutput, err error) TransactionCreator {
	return &transactionCreatorMock{
		output: output,
		err:    err,
	}
}

// Create mocks the Create behavior of TransactionInteractor
func (tcm transactionCreatorMock) Create(accountID, operationTypeID uint64, amount float64) (*presenter.CreateTransactionOutput, error) {
	if tcm.err != nil {
		return nil, tcm.err
	}

	return tcm.output, nil
}
