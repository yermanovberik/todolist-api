package main

import (
	"log"
	"net/http"
	"todo-list-api/internal/handler"
)

func main() {

	http.HandleFunc("/api/todo-list/tasks", handler.CreateTask)
	log.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
