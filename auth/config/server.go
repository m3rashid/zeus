package config

import (
	"common/server"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var Server server.Server

var db *gorm.DB
var redisClient *redis.Client

func GetDb() (*gorm.DB, error) {
	if db == nil {
		gormDb, err := Server.GetDb(redisClient)
		if err != nil {
			return nil, err
		}
		db = gormDb
	}

	return db, nil
}

func GetRedisClient() (*redis.Client, error) {
	if redisClient == nil {
		client := Server.GetRedisClient()
		redisClient = client
	}

	return redisClient, nil
}
