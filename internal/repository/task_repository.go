package repository

import (
	"sync"

	"github.com/iabdulzahid/anra_assessment/internal/model"
)

type TaskRepository struct {
	tasks map[string]model.Task
	mu    sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]model.Task),
	}
}

func (r *TaskRepository) Save(task model.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
}

func (r *TaskRepository) GetAll() []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]model.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}
