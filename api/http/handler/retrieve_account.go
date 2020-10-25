package handler

import (
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
	responseWriter := newResponseWriter(rw)

	if err != nil {
		h.logger.Println("invalid accountID:", err)

		output := map[string]string{"error": "Invalid accountID"}
		responseWriter.outputResponse(http.StatusBadRequest, output)
		return
	}

	output, err := h.accountRetriever.RetrieveByID(uint64(accountID))
	if err != nil {
		h.logger.Println("error retrieving account:", err)
		if errors.Is(err, repository.ErrRegisterNotFound) {
			responseWriter.outputResponse(http.StatusNotFound, nil)
			return
		}

		responseWriter.outputResponse(http.StatusInternalServerError, nil)
		return
	}

	responseWriter.outputResponse(http.StatusOK, output)
	return
}

type accountRetrieverMock struct {
	retrieveAccountOutput *presenter.RetrieveAccountOutput
	err                   error
}

// NewAccountRetrieverMock creates a new AccountRetriever implementation for unit tests
func NewAccountRetrieverMock(rao *presenter.RetrieveAccountOutput, err error) AccountRetriever {
	return &accountRetrieverMock{retrieveAccountOutput: rao, err: err}
}

// RetrieveByID mocks the RetrieveByID behavior of the AccountRetriever
func (arm accountRetrieverMock) RetrieveByID(accountID uint64) (*presenter.RetrieveAccountOutput, error) {
	if arm.err != nil {
		return nil, arm.err
	}

	return arm.retrieveAccountOutput, nil
}
