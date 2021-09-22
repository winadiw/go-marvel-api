package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// instance for accessing redis
var RDB *redis.Client

var ctx = context.Background()

func GetForKey(key string) (string, error) {
	val, err := RDB.Get(ctx, key).Result()
	return val, err
}

func SetForKey(key string, data string, expiry time.Duration) error {
	err := RDB.Set(ctx, key, data, expiry).Err()
	return err
}
