package main

import (
	"log"

	"github.com/winadiw/go-marvel-api/cache"
	"github.com/winadiw/go-marvel-api/config"
	"github.com/winadiw/go-marvel-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	cache.ConnectRedis()

	router.SetupRoutes(app)

	host := config.Config("HOST")
	log.Fatal(app.Listen(host))

}
