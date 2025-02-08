package routes

import (
	"api/database"
	"github.com/gofiber/fiber/v2"
)

// Road represents the road structure
type Road struct {
	ID     int      `json:"id" db:"id"`
	Name   string   `json:"name" db:"name"`
	Cities []string `json:"cities" db:"cities"`
}

// Hello sends a simple "Hello, World!" message when accessed via /hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// GetAllRoads handles the GET request to retrieve all roads
func GetAllRoads(c *fiber.Ctx) error {
	var roads []Road
	err := database.DB.Select(&roads, "SELECT * FROM roads")
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}
	return c.JSON(roads)
}

// CreateRoad handles the POST request to create a new road
func CreateRoad(c *fiber.Ctx) error {
	road := new(Road)
	if err := c.BodyParser(road); err != nil {
		return c.Status(400).SendString("Request error: " + err.Error())
	}

	// Insert road into the database
	_, err := database.DB.NamedExec(`INSERT INTO roads (name, cities) VALUES (:name, :cities)`, road)
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}

	return c.Status(201).JSON(road)
}
