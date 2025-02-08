package routes

import "github.com/gofiber/fiber/v2"

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Register routes for hello
	app.Get("/hello", Hello)
	app.Get("/getAllRoads", GetAllRoads)
}
