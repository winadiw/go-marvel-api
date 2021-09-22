package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/cache"
	"github.com/winadiw/go-marvel-api/external"
)

func GetCharacterById(c *fiber.Ctx) error {
	type GetCharactersByIdResponse struct {
		Id          int
		Name        string
		Description string
	}

	id := c.Params("id")

	marvelResponse, err := external.MarvelGetCharacterById(id)

	if err != nil {
		return c.Status(err.Code).JSON(err)
	}

	data := marvelResponse.Data.Results[0]

	response := GetCharactersByIdResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
	cache.CacheResponse(c, response)
	return c.JSON(response)
}
