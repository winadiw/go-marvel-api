package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Status handles api status
func Status(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello World!"})
}
