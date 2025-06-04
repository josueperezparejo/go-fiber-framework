package services

import (
	"github/josueperezparejo/my-go/config"
	"github/josueperezparejo/my-go/models"
)

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Return a list of all tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Find(&tasks)
	return tasks, result.Error
}

// GetTaskByID godoc
// @Summary Get a task by ID
// @Description Return a task based on its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID of the task"
// @Success 200 {object} models.Task
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

// GetTasksByProjectID godoc
// @Summary Get tasks by project ID
// @Description Return all tasks associated with a specific project
// @Tags tasks
// @Accept json
// @Produce json
// @Param project_id path int true "ID of the project"
// @Success 200 {array} models.Task
// @Failure 500 {object} map[string]string
// @Router /projects/{project_id}/tasks [get]
func GetTasksByProjectID(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Where("project_id = ?", projectID).Find(&tasks)
	return tasks, result.Error
}

// CreateTask godoc
// @Summary Create a task
// @Description Create a new task in the database
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body dtos.CreateTaskDTO true "Task to create"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task based on its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID of the task"
// @Param task body dtos.UpdateTaskDTO true "Data of the task to update"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func UpdateTask(id uint, updated *models.Task) error {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}
	task.Title = updated.Title
	task.Description = updated.Description
	task.Completed = updated.Completed
	task.ProjectID = updated.ProjectID
	return config.DB.Save(&task).Error
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task based on its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID of the task"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [delete]
func DeleteTask(id uint) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
