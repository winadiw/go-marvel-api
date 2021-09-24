package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winadiw/go-marvel-api/external"
	"github.com/winadiw/go-marvel-api/utils"
)

var (
	marvelGetCharacterById func(ID string) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData)
	marvelGetCharacters    func(limit, offset int) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData)
)

type marvelServiceMock struct{}

func (sm *marvelServiceMock) MarvelGetCharacterById(ID string) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
	return marvelGetCharacterById(ID)
}
func (sm *marvelServiceMock) MarvelGetCharacters(limit, offset int) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
	return marvelGetCharacters(limit, offset)
}

func TestGetCharacterById(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	marvelGetCharacterById = func(ID string) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
		response := external.MarvelGetCharactersResponse{
			BaseMarvelResponse: external.BaseMarvelResponse{},
		}
		response.Data.Results = []external.MarvelCharacterData{
			{
				ID:          1,
				Name:        "Zodiak",
				Description: "Test",
			},
		}
		return response, nil
	}
	app.Get("/api/characters/:id", GetCharacterById)
	req := httptest.NewRequest("GET", "/api/characters/1", nil)
	resp, _ := app.Test(req, 1)
	assert.Equalf(t, 200, resp.StatusCode, "Found Character")
}

func TestGetCharacterById_NotFound(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	marvelGetCharacterById = func(ID string) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
		return external.MarvelGetCharactersResponse{}, &utils.ResponseErrorData{
			Code:    404,
			Message: "We can't find that character",
			Status:  "error",
			Data:    nil,
		}
	}

	app.Get("/api/characters/:id", GetCharacterById)
	req := httptest.NewRequest("GET", "/api/characters/1", nil)
	resp, _ := app.Test(req, 1)
	assert.Equalf(t, 404, resp.StatusCode, "Character Not Found")
}
