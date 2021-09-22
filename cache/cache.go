package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// instance for accessing redis
var RDB *redis.Client

var ctx = context.Background()

func Get(key string) ([]byte, error) {
	return RDB.Get(ctx, key).Bytes()
}
func Set(key string, val string, ttl time.Duration) error {
	return RDB.Set(ctx, key, val, ttl).Err()
}

func Delete(key string) error {
	RDB.Del(ctx, key)
	return nil
}

func Reset() error {
	RDB.FlushAll(ctx)
	return nil
}

func Close() error {
	RDB.Close()
	return nil
}

func CacheResponse(c *fiber.Ctx, response interface{}) {
	key := c.Locals("cacheKey")
	if key == nil {
		fmt.Println("empty key")
		return
	}
	responseJson, errJson := json.Marshal(response)
	keyString := fmt.Sprintf("%s", key)

	if errJson != nil {
		fmt.Println(errJson)
		return
	}

	fmt.Println(errJson)
	err := Set(keyString, string(responseJson), 15*time.Minute)
	if err != nil {
		fmt.Println(err)
	}
}
