package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	app := fiber.New()
	app.Get("/api/status", Status)
	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "get HTTP status 200",
			route:        "/api/status",
			expectedCode: 200,
		},
	}
	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.name)
	}
}
