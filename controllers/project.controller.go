package controllers

import (
	"github/josueperezparejo/my-go/dtos"
	"github/josueperezparejo/my-go/models"
	"github/josueperezparejo/my-go/services"
	"github/josueperezparejo/my-go/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	projects, err := services.GetAllProjects()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error getting projects"})
	}
	return c.JSON(projects)
}

func GetProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	project, err := services.GetProjectByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Project not found"})
	}
	return c.JSON(project)
}

func CreateProject(c *fiber.Ctx) error {
	var dto dtos.CreateProjectDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	validationErrors := utils.ValidateStruct(dto)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	project := models.Project{
		Name:        dto.Name,
		Description: dto.Description,
	}

	if err := services.CreateProject(&project); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create project"})
	}
	return c.Status(201).JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {

	var dto dtos.UpdateProjectDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	validationErrors := utils.ValidateStruct(dto)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Project
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := services.UpdateProject(uint(id), &updated); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update project"})
	}
	return c.JSON(fiber.Map{"message": "Project updated"})
}

func DeleteProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := services.DeleteProject(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete project"})
	}
	return c.JSON(fiber.Map{"message": "Project deleted"})
}
