package main

// @title my API with Fiber
// @version 1.0
// @description This is a example API for a CRUD of projects and tasks.
// @host localhost:3000
// @BasePath /api
// @schemes http https

import (
	"github/josueperezparejo/my-go/config"
	"github/josueperezparejo/my-go/docs"
	"github/josueperezparejo/my-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{}))

	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Fiber app running with PostgreSQL!")
	})

	routes.SetupRoutes(app)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Listen(":3000")
}
