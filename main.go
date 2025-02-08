package main

import (
	"api/routes" // Import the routes package
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
