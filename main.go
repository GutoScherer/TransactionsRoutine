package main

import (
	"github.com/GutoScherer/TransactionsRoutine/api"
	"github.com/GutoScherer/TransactionsRoutine/api/http"
	"github.com/GutoScherer/TransactionsRoutine/infra/database/mysql"
	"github.com/gorilla/mux"
)

func main() {
	config := mysql.NewConfig("root", "admin123", "TransactionsRoutineDatabase", "3306", "transaction_routine")
	db, err := mysql.NewGormConnection(config)
	if err != nil {
		//panic(err)
	}

	router := mux.NewRouter()

	var httpServer api.Server = http.NewServer(router, db)
	httpServer.ListenAndServe(8080)
}
