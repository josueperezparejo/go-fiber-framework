package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	ProjectRoutes(api)
	TaskRoutes(api)
	ProjectTaskRoutes(api)
}
