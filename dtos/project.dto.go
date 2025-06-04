package dtos

type CreateProjectDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type UpdateProjectDTO struct {
	Name        string `json:"name" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,max=255"`
}
