package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"todo-list-api/internal/domain"
	"todo-list-api/internal/dto"
	"todo-list-api/internal/service"
)

var (
	tasks   = make(map[string]domain.Task)
	counter = 1
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO dto.TaskRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := service.CreateTask(taskDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO dto.TaskRequestDTO

	err := json.NewDecoder(r.Body).Decode(&taskDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	err = service.UpdateTask(id, taskDTO)

	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if err.Error() == "task title too long" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := service.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := service.MarkTaskDone(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
func ListTasks(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]
	if status == "" {
		status = r.URL.Query().Get("status")
	}

	if status == "" {
		status = "active"
	}

	filteredTasks, err := service.ListTasks(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(filteredTasks) == 0 {
		json.NewEncoder(w).Encode([]domain.Task{})
	} else {
		json.NewEncoder(w).Encode(filteredTasks)
	}
}
