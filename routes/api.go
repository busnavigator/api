package routes

import (
	"api/database"
	"github.com/gofiber/fiber/v2"
)

// Road represents the road structure
type Route struct {
	ID     int      `json:"id" db:"id"`
	Name   string   `json:"name" db:"name"`
	Cities []string `json:"cities" db:"cities"`
}

// Hello sends a simple "Hello, World!" message when accessed via /hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// GetAllRoads handles the GET request to retrieve all roads
func GetAllRoutes(c *fiber.Ctx) error {
	var routes []Route
	err := database.DB.Select(&routes, "SELECT * FROM routes")
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}
	return c.JSON(routes)
}

// CreateRoad handles the POST request to create a new road
func CreateRoad(c *fiber.Ctx) error {
	route := new(Route)
	if err := c.BodyParser(route); err != nil {
		return c.Status(400).SendString("Request error: " + err.Error())
	}

	// Insert road into the database
	_, err := database.DB.NamedExec(`INSERT INTO roads (name, cities) VALUES (:name, :cities)`, route)
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}

	return c.Status(201).JSON(route)
}
