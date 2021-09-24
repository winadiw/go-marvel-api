package middleware

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/winadiw/go-marvel-api/cache"
)

func stripQueryParam(inURL string, stripKey string) string {
	u, err := url.Parse(inURL)
	if err != nil {
		fmt.Println("Not an url: " + inURL)
		return inURL
	}
	q := u.Query()
	q.Del(stripKey)
	u.RawQuery = q.Encode()
	return u.String()
}

// getKey returns key to use for caching each request
func getKey(c *fiber.Ctx) string {
	return stripQueryParam(c.OriginalURL(), "refresh")
}

// EnableCache middleware to check cache if any
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
	result, err := cache.Get(key)

	// Store in locals for handler to set
	c.Locals("cacheKey", key)

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

	// Go to next middleware:
	return c.Next()
}
