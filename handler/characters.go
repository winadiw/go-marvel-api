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

	marvelResponse, err := external.MarvelGetCharacterById(id)

	if err != nil {
		return c.Status(err.Code).JSON(err)
	}

	data := marvelResponse.Data.Results[0]

	var response GetCharactersByIdResponse
	response.Id = int64(data.ID)
	response.Name = data.Name
	response.Description = data.Description

	return c.JSON(response)
}
