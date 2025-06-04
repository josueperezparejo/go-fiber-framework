package routes

import (
	"github/josueperezparejo/my-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(router fiber.Router) {
	task := router.Group("/tasks")
	task.Get("/", controllers.GetTasks)
	task.Get("/:id", controllers.GetTask)
	task.Post("/", controllers.CreateTask)
	task.Put("/:id", controllers.UpdateTask)
	task.Delete("/:id", controllers.DeleteTask)
}

func ProjectTaskRoutes(router fiber.Router) {
	projectTasks := router.Group("/projects/:projectId/tasks")
	projectTasks.Get("/", controllers.GetTasksByProject)
}
