package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharacterById(t *testing.T) {
	app.Get("/api/characters/:id", GetCharacterById)
	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "Get HTTP status 200",
			route:        "/api/characters/1011127",
			expectedCode: 200,
		},
	}
	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.name)
	}
}
