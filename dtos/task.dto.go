package dtos

type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ProjectID   uint   `json:"project_id" binding:"required"`
}

type UpdateTaskDTO struct {
	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,max=255"`
	Completed   bool   `json:"completed"`
}
