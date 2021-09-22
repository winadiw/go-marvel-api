package cache

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/winadiw/go-marvel-api/config"
)

// ConnectRedis connects to redis client
func ConnectRedis() {
	configRedisDb, err := strconv.Atoi(config.Config("REDIS_DB"))
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Config("REDIS_ADDR"),
		Password: config.Config("REDIS_PASS"), // no password set
		DB:       configRedisDb,               // use default DB
	})

	log.Println("Connected to Redis")
}
