package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/cache"
)

func getKey(c *fiber.Ctx) string {
	return c.Path()
}

func EnableCache(c *fiber.Ctx) error {
	// Only cache GET method
	if c.Method() != fiber.MethodGet {
		return c.Next()
	}

	// If need to refresh, call next
	if c.Query("refresh") == "true" {
		return c.Next()
	}

	key := getKey(c)

	fmt.Println(key)

	result, err := cache.Get(key)

	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		// means have cache
		var response interface{}

		json.Unmarshal(result, &response)
		return c.JSON(response)
	}

	c.Locals("cacheKey", key)

	// Go to next middleware:
	return c.Next()
}
