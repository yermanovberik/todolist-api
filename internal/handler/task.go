package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todo-list-api/internal/domain"
)

var tasks = make(map[string]domain.Task)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = uuid.New().String()
	task.ActiveAt, _ = time.Parse("2006-01-02", task.ActiveAt.String())
	tasks[strconv.Itoa(task.ID)] = task

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": task.ID})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Реализация логики обновления задачи
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Реализация логики удаления задачи
}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	// Реализация логики пометить задачу выполненной
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	// Реализация логики получения списка задач
}
