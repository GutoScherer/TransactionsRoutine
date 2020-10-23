package main

import (
	"log"
	"os"

	"github.com/GutoScherer/TransactionsRoutine/api"
	"github.com/GutoScherer/TransactionsRoutine/api/http"
	"github.com/GutoScherer/TransactionsRoutine/infra/database/mysql"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("Initiating app")

	config := mysql.NewConfig(
		os.Getenv(`MYSQL_USER`),
		os.Getenv(`MYSQL_PASSWORD`),
		os.Getenv(`MYSQL_HOST`),
		os.Getenv(`MYSQL_PORT`),
		os.Getenv(`MYSQL_DATABASE`),
	)
	db, err := mysql.NewGormConnection(config)
	if err != nil {
		logger.Fatalln("error connecting to database:", err.Error())
		return
	}

	router := mux.NewRouter()

	var httpServer api.Server = http.NewServer(router, db, logger)
	httpServer.ListenAndServe(8080)
}
