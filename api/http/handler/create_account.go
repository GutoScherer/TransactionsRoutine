package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
)

type createAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

// AccountCreator represents an creator for account usecases
type AccountCreator interface {
	Create(documentNumber string) (*presenter.CreateAccountOutput, error)
}

// CreateAccount represents the http handler struct to create an account
type CreateAccount struct {
	accountCreator AccountCreator
	logger         *log.Logger
}

// NewCreateAccount creates a new CreateAccount struct
func NewCreateAccount(ac AccountCreator, logger *log.Logger) *CreateAccount {
	return &CreateAccount{
		accountCreator: ac,
		logger:         logger,
	}
}

// HandlerFunc is the http handler function used by the server
func (h CreateAccount) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	var createAccRequest createAccountRequest
	responseWriter := newResponseWriter(rw)

	err := json.NewDecoder(r.Body).Decode(&createAccRequest)
	if err != nil {
		h.logger.Println("invalid request body:", err)

		output := map[string]string{"error": "invalid JSON body"}
		responseWriter.outputResponse(http.StatusBadRequest, output)
		return
	}

	output, err := h.accountCreator.Create(createAccRequest.DocumentNumber)

	if err != nil {
		h.logger.Println("error creating account:", err)

		if errors.Is(err, repository.ErrDuplicatedEntry) {
			responseWriter.outputResponse(http.StatusConflict, nil)
			return
		}

		if errors.Is(err, repository.ErrInvalidData) {
			responseWriter.outputResponse(http.StatusUnprocessableEntity, nil)
			return
		}

		responseWriter.outputResponse(http.StatusInternalServerError, nil)
		return
	}

	responseWriter.outputResponse(http.StatusInternalServerError, output)
	return
}
