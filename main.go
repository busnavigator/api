package main

import (
	"github.com/gofiber/fiber/v2"
	"api/routes" // Import the routes package
)

func main() {
	app := fiber.New()

	// Initialize routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
