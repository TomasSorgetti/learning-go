package services

import (
    "errors"
    "go-app/internal/models"
    "go-app/internal/repository"
)

type TaskService struct {
    repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) Create(task *models.Task) error {
    if task.Title == "" {
        return errors.New("title cannot be empty")
    }
    return s.repo.Create(task)
}

func (s *TaskService) GetAll() ([]models.Task, error) {
    return s.repo.GetAll()
}

func (s *TaskService) GetByID(id int) (*models.Task, error) {
    return s.repo.GetByID(id)
}

func (s *TaskService) Update(task *models.Task) error {
    if task.Title == "" {
        return errors.New("title cannot be empty")
    }
    return s.repo.Update(task)
}

func (s *TaskService) Delete(id int) error {
    return s.repo.Delete(id)
}