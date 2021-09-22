package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	fibercache "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/winadiw/go-marvel-api/cache"
	"github.com/winadiw/go-marvel-api/handler"
	"github.com/winadiw/go-marvel-api/utils"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(requestid.New())
	app.Use(fibercache.New(fibercache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
		Storage:      cache.RedisCaching{},
	}))
	api := app.Group("/api", logger.New(logger.Config{
		Format:     utils.LoggerFormat(),
		TimeFormat: "2006-01-02T15:04:05Z07:00",
		TimeZone:   "UTC",
	})) //Logger will affect all API with prefix /api***

	api.Get("/status", handler.Status)

	// Characters
	characters := api.Group("/characters")
	characters.Get("/:id", handler.GetCharactersById)
}
