package utils

import "github.com/gofiber/fiber/v2"

func ResponseError(message string, data interface{}) map[string]interface{} {
	return fiber.Map{
		"status": "error", "message": message, "data": data}
}
