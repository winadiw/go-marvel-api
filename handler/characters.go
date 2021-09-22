package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/external"
)

func GetCharactersById(c *fiber.Ctx) error {
	type GetCharactersByIdResponse struct {
		Id          int64
		Name        string
		Description string
	}

	id := c.Params("id")

	var response GetCharactersByIdResponse

	external.MarvelGetCharactersById(id)

	return c.JSON(response)
}
