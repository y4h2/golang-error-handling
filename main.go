package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/y4h2/golang-error-handling/app/api"
	"github.com/y4h2/golang-error-handling/app/repository"
	"github.com/y4h2/golang-error-handling/app/service"
)

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlers := api.New(service.New(repository.New(db)))
	http.HandleFunc("/article", handlers.Get)
	http.ListenAndServe(":9000", nil)
}
