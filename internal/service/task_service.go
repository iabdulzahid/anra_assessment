package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/iabdulzahid/anra_assessment/internal/model"
	"github.com/iabdulzahid/anra_assessment/internal/repository"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

var validStatuses = map[string]bool{
	StatusTodo:       true,
	StatusInProgress: true,
	StatusDone:       true,
}

// TaskService contains business logic for tasks
type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string, status string) (model.Task, error) {

	if title == "" {
		return model.Task{}, errors.New("title is required")
	}

	if len(title) > 200 {
		return model.Task{}, errors.New("title cannot exceed 200 characters")
	}

	if status == "" {
		status = StatusTodo
	}

	if !validStatuses[status] {
		return model.Task{}, errors.New("invalid status")
	}

	task := model.Task{
		ID:     uuid.New().String(),
		Title:  title,
		Status: status,
	}

	s.repo.Save(task)

	return task, nil
}

func (s *TaskService) ListTasks() []model.Task {
	return s.repo.GetAll()
}
