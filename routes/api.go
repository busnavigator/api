package routes

import (
	"api/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

// Route represents the route structure
type Route struct {
	ID       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	NextStop int      `json:"nextStop" db:"nextStop"`
	Stops    []string `json:"stops" db:"stops"`
}

// Hello sends a simple "Hello, World!" message when accessed via /hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// GetAllRoutes handles the GET request to retrieve all routes
func GetAllRoutes(c *fiber.Ctx) error {
	var routes []Route
	rows, err := database.DB.Query("SELECT id, name, nextStop, stops FROM routes")
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var route Route
		var stopsJSON []byte
		if err := rows.Scan(&route.ID, &route.Name, &route.NextStop, &stopsJSON); err != nil {
			return c.Status(500).SendString("Database scan error: " + err.Error())
		}
		if err := json.Unmarshal(stopsJSON, &route.Stops); err != nil {
			return c.Status(500).SendString("JSON decoding error: " + err.Error())
		}
		routes = append(routes, route)
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
	_, err = database.DB.Exec(`INSERT INTO routes (name, nextStop, stops) VALUES ($1, $2, $3)`, route.Name, route.NextStop, stopsJSON)
	if err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}

	return c.Status(201).JSON(route)
}
