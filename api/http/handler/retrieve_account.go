package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/gorilla/mux"
)

// AccountRetriever represents an retriever for account usecases
type AccountRetriever interface {
	RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error)
}

// RetrieveAccount represents the http handler struct to retrieve an account
type RetrieveAccount struct {
	accountRetriever AccountRetriever
}

// NewRetrieveAccount creates a new RetrieveAccount struct
func NewRetrieveAccount(ar AccountRetriever) *RetrieveAccount {
	return &RetrieveAccount{accountRetriever: ar}
}

// HandlerFunc is the http handler function used by the server
func (h RetrieveAccount) HandlerFunc(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["accountID"])
	if err != nil {
		// TODO: Error handling
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	output, err := h.accountRetriever.RetrieveByID(uint64(accountID))
	if err != nil {
		// TODO: Error handling
		http.Error(rw, "", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(output)
}
