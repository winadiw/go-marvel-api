package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/cache"
	"github.com/winadiw/go-marvel-api/external"
	"github.com/winadiw/go-marvel-api/utils"
)

// GetCharacterById handles incoming request for get character by id
func GetCharacterById(c *fiber.Ctx) error {
	type GetCharactersByIdResponse struct {
		Id          int
		Name        string
		Description string
	}

	id := c.Params("id")

	marvelResponse, err := external.MarvelService.MarvelGetCharacterById(id)

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

// GetCharacterList handles incoming request for get characters
func GetCharacterList(c *fiber.Ctx) error {
	limit := c.Query("limit", "100")
	offset := c.Query("offset", "0")

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError(fiber.StatusBadRequest,
			"Unable to parse limit to int", nil))
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError(fiber.StatusBadRequest,
			"Unable to parse offset to int", nil))
	}

	marvelResponse, errMarvel := external.MarvelService.MarvelGetCharacters(limitInt, offsetInt)

	if errMarvel != nil {
		return c.Status(errMarvel.Code).JSON(errMarvel)
	}

	response := marvelResponse.IdList()

	cache.CacheResponse(c, response)

	return c.JSON(response)
}
