package routes

import (
	"github/josueperezparejo/my-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	project := router.Group("/projects")
	project.Get("/", controllers.GetProjects)
	project.Get("/:id", controllers.GetProject)
	project.Post("/", controllers.CreateProject)
	project.Put("/:id", controllers.UpdateProject)
	project.Delete("/:id", controllers.DeleteProject)
}
