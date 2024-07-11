package service

import (
	"errors"
	"sort"
	"time"
	"todo-list-api/internal/domain"
	"todo-list-api/internal/dto"
	"todo-list-api/internal/repository"
)

func CreateTask(taskDTO dto.TaskRequestDTO) (domain.Task, error) {
	if len(taskDTO.Title) > 200 {
		return domain.Task{}, errors.New("task title too long")
	}

	activeAt, err := time.Parse("2006-01-02", taskDTO.ActiveAt)

	if err != nil {
		return domain.Task{}, err
	}

	task, err := repository.FindTaskByTitleAndDate(taskDTO.Title, activeAt)
	if err == nil {
		return domain.Task{}, errors.New("task already exists")
	}

	task = domain.Task{
		ID:       repository.GetNextTaskID(),
		Tittle:   taskDTO.Title,
		ActiveAt: activeAt,
		Done:     false,
	}

	repository.SaveTask(task)
	return task, nil
}

func UpdateTask(id string, taskDto dto.TaskRequestDTO) error {
	task, err := repository.FindTaskByID(id)
	if err != nil {
		return errors.New("task not found")
	}

	if len(taskDto.Title) > 200 {
		return errors.New("task title too long")
	}
	task.Tittle = taskDto.Title
	activeAt, err := time.Parse("2006-01-02", taskDto.ActiveAt)
	if err != nil {
		return err
	}

	task.ActiveAt = activeAt
	err = repository.SaveTask(task)
	if err != nil {
		return err
	}
	return nil
}

func MarkTaskDone(id string) error {
	task, err := repository.FindTaskByID(id)
	if err != nil {
		return errors.New("task not found")
	}

	task.Done = true
	err = repository.SaveTask(task)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	err := repository.DeleteTask(id)
	if err != nil {
		return errors.New("task not found")
	}
	return nil
}

func ListTasks(status string) ([]domain.Task, error) {
	currentTime := time.Now()
	tasks, err := repository.ListTasks()
	if err != nil {
		return nil, errors.New("error retrieving tasks")
	}

	var filteredTasks []domain.Task

	if status == "done" {
		for _, task := range tasks {
			if task.Done {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		for _, task := range tasks {
			if !task.Done && (task.ActiveAt.Before(currentTime) || task.ActiveAt.Equal(currentTime)) {
				if task.ActiveAt.Weekday() == time.Saturday || task.ActiveAt.Weekday() == time.Sunday {
					task.Tittle = "ВЫХОДНОЙ - " + task.Tittle
				}
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	sort.Slice(filteredTasks, func(i, j int) bool {
		return filteredTasks[i].ID < filteredTasks[j].ID
	})

	return filteredTasks, nil
}
