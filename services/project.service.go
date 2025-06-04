package services

import (
	"github/josueperezparejo/my-go/config"
	"github/josueperezparejo/my-go/models"
)

// GetAllProjects godoc
// @Summary Get all projects
// @Description Return a list of all projects
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} map[string]string
// @Router /projects [get]
func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	result := config.DB.Preload("Tasks").Find(&projects)
	return projects, result.Error
}

// GetProjectByID godoc
// @Summary Get a project by ID
// @Description Return a project based on its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "ID of the project"
// @Success 200 {object} models.Project
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [get]
func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	result := config.DB.Preload("Tasks").First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

// CreateProject godoc
// @Summary Create a project
// @Description Create a new project in the database
// @Tags projects
// @Accept json
// @Produce json
// @Param project body dtos.CreateProjectDTO true "Project to create"
// @Success 201 {object} models.Project
// @Failure 400 {object} map[string]string
// @Router /projects [post]
func CreateProject(project *models.Project) error {
	return config.DB.Create(project).Error
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update an existing project based on its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "ID of the project"
// @Param project body dtos.UpdateProjectDTO true "Project to update"
// @Success 200 {object} models.Project
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [put]
func UpdateProject(id uint, updated *models.Project) error {
	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		return err
	}
	project.Name = updated.Name
	project.Description = updated.Description
	return config.DB.Save(&project).Error
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project based on its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "ID of the project"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [delete]
func DeleteProject(id uint) error {
	return config.DB.Delete(&models.Project{}, id).Error
}
