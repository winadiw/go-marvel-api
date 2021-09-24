package integration__test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/winadiw/go-marvel-api/handler"
)

// Integration Test
func TestIntegrationMarvelGetCharacterById(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for integration to marvel get character by id")
	}
	app := fiber.New()
	app.Get("/api/characters/:id", handler.GetCharacterById)
	req := httptest.NewRequest("GET", "/api/characters/1011127", nil)
	resp, err := app.Test(req, -1)

	assert.Nil(t, err)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody interface{}
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 200, resp.StatusCode, "Found Character")
}

func TestIntegrationMarvelGetCharacters(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for integration to marvel get character id list")
	}
	app := fiber.New()
	app.Get("/api/characters", handler.GetCharacterList)
	req := httptest.NewRequest("GET", "/api/characters", nil)
	resp, err := app.Test(req, -1)

	assert.Nil(t, err)

	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody interface{}
	err = json.Unmarshal(responseData, &responseBody)
	assert.Nil(t, err)

	assert.Equalf(t, 200, resp.StatusCode, "Found Characters")
}
