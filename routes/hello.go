package routes

import "github.com/gofiber/fiber/v2"

// Hello handler
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
