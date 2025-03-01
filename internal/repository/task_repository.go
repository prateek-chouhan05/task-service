package repository

import (
	"task-service/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetAll(limit, offset int, status string) ([]models.Task, error) {
	var tasks []models.Task

	query := r.db.Limit(limit).Offset(offset)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Find(&tasks).Error
	return tasks, err
}

// CountTasks returns the total number of tasks (with optional filtering)
func (r *TaskRepository) CountTasks(status string) (int, error) {
	var count int64

	query := r.db.Model(&models.Task{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&count).Error
	return int(count), err
}


func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
