package main

import (
	"log"
	"net/http"
	"todo-list-api/internal/app"
)

func main() {
	router := app.NewRouter()

	log.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
