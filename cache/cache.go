package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// instance for accessing redis
var RDB *redis.Client

var ctx = context.Background()

// RedisCaching implements fiber.Storage
type RedisCaching struct {
}

func (rc RedisCaching) Get(key string) ([]byte, error) {
	return RDB.Get(ctx, key).Bytes()
}

func (rc RedisCaching) Set(key string, val []byte, ttl time.Duration) error {
	return RDB.Set(ctx, key, val, ttl).Err()
}

func (rc RedisCaching) Delete(key string) error {
	RDB.Del(ctx, key)
	return nil
}

func (rc RedisCaching) Reset() error {
	RDB.FlushAll(ctx)
	return nil
}

func (rc RedisCaching) Close() error {
	RDB.Close()
	return nil
}
