package services

import (
	"task-service/internal/models"
	"task-service/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetTasks(limit, offset int, status string) ([]models.Task, int, error) {
	tasks, err := s.repo.GetAll(limit, offset, status)
	if err != nil {
		return nil, 0, err
	}

	// Get total count of tasks
	totalCount, err := s.repo.CountTasks(status)
	if err != nil {
		return nil, 0, err
	}

	return tasks, totalCount, nil
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
