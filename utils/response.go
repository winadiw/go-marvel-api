package utils

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return fiber.Map{
		"status": "success", "message": message, "data": data}
}

func ResponseListSuccess(message string, data interface{}, items int, totalItems int64) map[string]interface{} {
	return fiber.Map{
		"status": "success", "message": message, "data": data,
		"items": items, "totalItems": totalItems}
}

func ResponseError(message string, data interface{}) map[string]interface{} {
	return fiber.Map{
		"status": "error", "message": message, "data": data}
}
