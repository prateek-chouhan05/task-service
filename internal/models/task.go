package models

import (
	"time"

	"gorm.io/gorm"
)

// StatusEnum defines allowed statuses
type StatusEnum string

const (
	Pending   StatusEnum = "Pending"
	InProgress StatusEnum = "InProgress"
	Completed  StatusEnum = "Completed"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      StatusEnum     `json:"status" gorm:"type:varchar(20);not null;check(status IN ('Pending', 'InProgress', 'Completed'))"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Status == "" {
		t.Status = Pending
	}
	return nil
}
