package server

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (s Server) GetRedisClient() *redis.Client {
	fmt.Println("Creating new redis client")
	return redis.NewClient(&redis.Options{Addr: s.RedisUri})
}
