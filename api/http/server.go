package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/infra/repository"
	"github.com/GutoScherer/TransactionsRoutine/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Server represents an HTTP server
type Server struct {
	router *mux.Router
	db     *gorm.DB
}

// NewServer creates a new Server struct
func NewServer(router *mux.Router, db *gorm.DB) *Server {
	return &Server{router: router, db: db}
}

// ListenAndServe listens on the specified port and exposes the server
func (s Server) ListenAndServe(port int) {
	s.router.HandleFunc("/accounts/{accountID}", s.retrieveAccountInfoHandler()).Methods("GET")
	s.router.HandleFunc("/accounts", s.createAccountHandler()).Methods("POST")
	s.router.HandleFunc("/transactions", s.createTransactionHandler()).Methods("POST")
	http.ListenAndServe(":8080", s.router)
}

func (s Server) retrieveAccountInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewAccountRepository(s.db)
		vars := mux.Vars(r)
		accountID, err := strconv.Atoi(vars["accountID"])
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		uc := usecase.NewRetrieveAccount(repo)
		acc, err := uc.RetrieveByID(uint64(accountID))
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(acc)
	}
}

func (s Server) createAccountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createAccRequest createAccountRequest

		err := json.NewDecoder(r.Body).Decode(&createAccRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		acc := entity.NewAccount(createAccRequest.DocumentNumber)

		repo := repository.NewAccountRepository(s.db)
		acc, err = repo.Save(acc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(acc)
	}
}

func (s Server) createTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createTransactionRequest createTransactionRequest

		err := json.NewDecoder(r.Body).Decode(&createTransactionRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		transaction, err := entity.NewTransaction(createTransactionRequest.AccountID, createTransactionRequest.OperationTypeID, createTransactionRequest.Amount)
		if err != nil {
			//TODO: Tipo de erro
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		repo := repository.NewTransactionRepository(s.db)
		transaction, err = repo.Save(transaction)
		if err != nil {
			//TODO: Tipo de erro
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transaction)
	}
}

type createAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type createTransactionRequest struct {
	AccountID       uint64  `json:"account_id"`
	OperationTypeID uint64  `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
