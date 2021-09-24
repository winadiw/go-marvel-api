package router

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestSetupRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	fiberApp := fiber.New()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Called",
			args: args{
				app: fiberApp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupRoutes(tt.args.app)
		})
	}
}
