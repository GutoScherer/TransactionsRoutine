package http

import (
	"net/http"

	"github.com/GutoScherer/TransactionsRoutine/api/http/handler"
	"github.com/GutoScherer/TransactionsRoutine/usecase/interactor"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
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
	repo := repository.NewAccountRepository(s.db)
	presenter := presenter.NewAccountPresenter()
	interactor := interactor.NewAccountInteractor(repo, presenter)

	handler := handler.NewRetrieveAccount(interactor)
	return handler.HandlerFunc
}

func (s Server) createAccountHandler() http.HandlerFunc {
	repo := repository.NewAccountRepository(s.db)
	presenter := presenter.NewAccountPresenter()
	interactor := interactor.NewAccountInteractor(repo, presenter)

	handler := handler.NewCreateAccount(interactor)
	return handler.HandlerFunc
}

func (s Server) createTransactionHandler() http.HandlerFunc {
	repo := repository.NewTransactionRepository(s.db)
	presenter := presenter.NewTransactionPresenter()
	interactor := interactor.NewTransactionInteractor(repo, presenter)

	handler := handler.NewCreateTransaction(interactor)
	return handler.HandlerFunc
}
