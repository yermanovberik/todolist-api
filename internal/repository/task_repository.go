package repository

import (
	"errors"
	"strconv"
	"sync"
	"time"
	"todo-list-api/internal/domain"
)

var (
	tasks   = make(map[string]domain.Task)
	counter = 1
	mu      sync.Mutex
)

func GetNextTaskID() int {
	mu.Lock()
	defer mu.Unlock()
	id := counter
	counter++
	return id
}

func FindTaskByID(id string) (domain.Task, error) {
	task, ok := tasks[id]
	if !ok {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		return nil
	}
	return errors.New("task not found")
}

func SaveTask(task domain.Task) error {
	mu.Lock()
	defer mu.Unlock()
	tasks[strconv.Itoa(task.ID)] = task
	return nil
}

func FindTaskByTitleAndDate(title string, date time.Time) (domain.Task, error) {
	for _, task := range tasks {
		if task.Tittle == title && task.ActiveAt.Equal(date) {
			return task, nil
		}
	}
	return domain.Task{}, errors.New("task not found")
}

func ListTasks() ([]domain.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	var allTasks []domain.Task
	for _, task := range tasks {
		allTasks = append(allTasks, task)
	}

	if len(allTasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return allTasks, nil
}
