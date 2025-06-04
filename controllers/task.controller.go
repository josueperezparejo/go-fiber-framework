package controllers

import (
	"github/josueperezparejo/my-go/dtos"
	"github/josueperezparejo/my-go/models"
	"github/josueperezparejo/my-go/services"
	"github/josueperezparejo/my-go/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks, err := services.GetAllTasks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error getting tasks"})
	}
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := services.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func GetTasksByProject(c *fiber.Ctx) error {
	projectId, _ := strconv.Atoi(c.Params("projectId"))
	tasks, err := services.GetTasksByProjectID(uint(projectId))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error getting tasks for project"})
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	var dto dtos.CreateTaskDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	validationErrors := utils.ValidateStruct(dto)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	task := models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   dto.Completed,
		ProjectID:   dto.ProjectID,
	}

	if err := services.CreateTask(&task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(201).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	var dto dtos.UpdateTaskDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	validationErrors := utils.ValidateStruct(dto)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	if err := services.UpdateTask(uint(id), &models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   dto.Completed,
	}); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(fiber.Map{"message": "Task updated"})
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := services.DeleteTask(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
	}
	return c.JSON(fiber.Map{"message": "Task deleted"})
}
