package utils

import "github.com/gofiber/fiber/v2"

func ResponseError(code int, message string, data interface{}) map[string]interface{} {
	return fiber.Map{
		"code":   code,
		"status": "error", "message": message, "data": data}
}
