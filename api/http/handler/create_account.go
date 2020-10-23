package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
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
}

// NewCreateAccount creates a new CreateAccount struct
func NewCreateAccount(ac AccountCreator) *CreateAccount {
	return &CreateAccount{accountCreator: ac}
}

// HandlerFunc is the http handler function used by the server
func (h CreateAccount) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	var createAccRequest createAccountRequest

	err := json.NewDecoder(r.Body).Decode(&createAccRequest)
	if err != nil {
		// TODO: Error handling
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.accountCreator.Create(createAccRequest.DocumentNumber)

	if err != nil {
		// TODO: Error handling
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(output)
}
