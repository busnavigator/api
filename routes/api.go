package routes

import (
	"api/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

// Road represents the road structure
type Route struct {
	ID    int      `json:"id" db:"id"`
	Name  string   `json:"name" db:"name"`
	Stops []string `json:"stops" db:"stops"`
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

func CreateRoute(c *fiber.Ctx) error {
	route := new(Route)
	if err := c.BodyParser(route); err != nil {
		return c.Status(400).SendString("Request error: " + err.Error())
	}

	// Convert stops slice to JSON
	stopsJSON, err := json.Marshal(route.Stops)
	if err != nil {
		return c.Status(500).SendString("JSON encoding error: " + err.Error())
	}

	// Insert into PostgreSQL (assuming `stops` column is of type JSONB)
	_, err = database.DB.Exec(`INSERT INTO routes (name, stops) VALUES ($1, $2)`, route.Name, stopsJSON)
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}

	return c.Status(201).JSON(route)
}
