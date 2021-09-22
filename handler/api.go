package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/utils"
)

// Status handles api status
func Status(c *fiber.Ctx) error {
	return c.JSON(utils.ResponseSuccess("Hello World", nil))
}
