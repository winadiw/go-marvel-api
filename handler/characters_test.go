package handler

import (
	"encoding/json"
	"io/ioutil"
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

func TestMarvelGetCharacterById(t *testing.T) {
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

func TestMarvelGetCharacterById_NotFound(t *testing.T) {
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

func TestMarvelGetCharacters(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	marvelGetCharacters = func(limit, offset int) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
		response := external.MarvelGetCharactersResponse{
			BaseMarvelResponse: external.BaseMarvelResponse{},
		}
		response.Data.Results = []external.MarvelCharacterData{
			{
				ID:          1,
				Name:        "Zodiak",
				Description: "Test",
			},
			{
				ID:          2,
				Name:        "Zodiak",
				Description: "Test",
			},
		}
		return response, nil
	}
	app.Get("/api/characters", GetCharacterList)
	req := httptest.NewRequest("GET", "/api/characters?limit=50&offset=0", nil)
	resp, _ := app.Test(req, 1)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody []int
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 200, resp.StatusCode, "Found Characters")
	assert.EqualValues(t, 1, responseBody[0], "Body is correct value for index 0")
	assert.EqualValues(t, 2, responseBody[1], "Body is correct value for index 1")
}

func TestMarvelGetCharacterList_ErrorNetwork(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	marvelGetCharacters = func(limit, offset int) (external.MarvelGetCharactersResponse, *utils.ResponseErrorData) {
		return external.MarvelGetCharactersResponse{}, &utils.ResponseErrorData{
			Code:    422,
			Message: "Network Error",
			Status:  "error",
			Data:    nil,
		}
	}
	app.Get("/api/characters", GetCharacterList)
	req := httptest.NewRequest("GET", "/api/characters?limit=50&offset=0", nil)
	resp, _ := app.Test(req, 1)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody utils.ResponseErrorData
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 422, resp.StatusCode, "Network Error")
	assert.EqualValues(t, 422, responseBody.Code, "Code is correct")
	assert.EqualValues(t, "Network Error", responseBody.Message, "Message is correct")
}

func TestMarvelGetCharacterList_LimitNotInteger(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	app.Get("/api/characters", GetCharacterList)
	req := httptest.NewRequest("GET", "/api/characters?limit=abcd&offset=0", nil)
	resp, _ := app.Test(req, 1)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody utils.ResponseErrorData
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 400, resp.StatusCode, "Error limit is not a number")
	assert.Equalf(t, 400, responseBody.Code, "Error Body Code")
	assert.Equalf(t, "Unable to parse limit to int", responseBody.Message, "Error Body Message")
}

func TestMarvelGetCharacterList_OffsetNotInteger(t *testing.T) {
	external.MarvelService = &marvelServiceMock{}
	app.Get("/api/characters", GetCharacterList)
	req := httptest.NewRequest("GET", "/api/characters?limit=0&offset=bcde", nil)
	resp, _ := app.Test(req, 1)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody utils.ResponseErrorData
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 400, resp.StatusCode, "Error offset is not a number")
	assert.Equalf(t, 400, responseBody.Code, "Error Body Code")
	assert.Equalf(t, "Unable to parse offset to int", responseBody.Message, "Error Body Message")
}
