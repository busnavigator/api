package main

import (
	"api/database"
	"api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDB()

	// Setup API
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
