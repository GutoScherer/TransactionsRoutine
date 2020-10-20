package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", createAccountHandler).Methods("POST")

	router.HandleFunc("/accounts/{accountId}", retrieveAccountInfoHandler).Methods("GET")

	router.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: Create account route"))
}

func retrieveAccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: Retrieve account informations route"))
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: Create transaction route"))
}
