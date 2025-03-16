package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"hello-world/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, or specify the frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Logger
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${ip}\n",
	}))

	// Set up routes
	routes.SetupRoutes(app)

	// Start server
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
