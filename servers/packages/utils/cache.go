package utils

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient(redisUri string) *redis.Client {
	fmt.Println("Creating new redis client")
	return redis.NewClient(&redis.Options{Addr: redisUri})
}
