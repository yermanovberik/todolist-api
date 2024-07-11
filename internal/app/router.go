package app

import (
	"github.com/gorilla/mux"
	"todo-list-api/internal/handler"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todo-list/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/api/todo-list/list-tasks/{status}", handler.ListTasks).Methods("GET")
	router.HandleFunc("/api/todo-list/mark-task/{id}", handler.MarkTaskDone).Methods("PUT")
	router.HandleFunc("/api/todo-list/task/{id}", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/todo-list/tasks/delete/{id}", handler.DeleteTask).Methods("DELETE")

	return router
}
