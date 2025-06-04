package models

import (
	"time"

	"gorm.io/gorm"
)

// @Description Task model with GORM embedded fields
type Task struct {
	// @Description The unique identifier of the task
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	// @Description The creation timestamp of the task
	CreatedAt time.Time `json:"created_at" example:"2024-06-04T12:00:00Z"`
	// @Description The last update timestamp of the task
	UpdatedAt time.Time `json:"updated_at" example:"2024-06-04T12:00:00Z"`
	// @Description The deletion timestamp of the task (if deleted)
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" example:"null"`
	// @Description The title of the task
	Title string `json:"title" gorm:"not null" example:"Complete project documentation"`
	// @Description The description of the task
	Description string `json:"description" example:"Write API documentation using Swagger"`
	// @Description The completion status of the task
	Completed bool `json:"completed" gorm:"default:false" example:"false"`
	// @Description The ID of the project this task belongs to
	ProjectID uint `json:"project_id" gorm:"not null" example:"1"`
}
