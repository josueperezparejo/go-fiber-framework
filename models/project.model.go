package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	// @Description The unique identifier of the project
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	// @Description The creation timestamp of the project
	CreatedAt time.Time `json:"created_at" example:"2024-06-04T12:00:00Z"`
	// @Description The last update timestamp of the project
	UpdatedAt time.Time `json:"updated_at" example:"2024-06-04T12:00:00Z"`
	// @Description The deletion timestamp of the project (if deleted)
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" example:"null"`
	// @Description The name of the project
	Name string `json:"name" gorm:"not null" example:"Project 1"`
	// @Description The description of the project
	Description string `json:"description" example:"This is a description of the project"`
	// @Description The tasks associated with the project
	Tasks []Task `json:"tasks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
