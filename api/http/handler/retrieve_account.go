package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
	"github.com/gorilla/mux"
)

// AccountRetriever represents an retriever for account usecases
type AccountRetriever interface {
	RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error)
}

// RetrieveAccount represents the http handler struct to retrieve an account
type RetrieveAccount struct {
	accountRetriever AccountRetriever
	logger           *log.Logger
}

// NewRetrieveAccount creates a new RetrieveAccount struct
func NewRetrieveAccount(ar AccountRetriever, logger *log.Logger) *RetrieveAccount {
	return &RetrieveAccount{
		accountRetriever: ar,
		logger:           logger,
	}
}

// HandlerFunc is the http handler function used by the server
func (h RetrieveAccount) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["accountID"])
	if err != nil {
		h.logger.Println("invalid accountID:", err)

		output := map[string]string{"error": "Invalid accountID"}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(output)
		return
	}

	output, err := h.accountRetriever.RetrieveByID(uint64(accountID))
	if err != nil {
		h.logger.Println("error retrieving account:", err)
		if errors.Is(err, repository.ErrRegisterNotFound) {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(output)
	return
}
