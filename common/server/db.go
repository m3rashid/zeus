package server

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"time"

	"github.com/Pacific73/gorm-cache/cache"
	"github.com/Pacific73/gorm-cache/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDBaseModel struct {
	ID uint `gorm:"primary_key;column:id" json:"id"`
}

type DefaultBaseModel struct {
	IDBaseModel
	Deleted   bool      `gorm:"column:deleted;default:false" json:",omitempty" validate:""`
	CreatedAt time.Time `gorm:"column:createdAt; default:current_timestamp" json:"createdAt"`
}

func (server *Server) GetDb(redisClient *redis.Client) (*gorm.DB, error) {
	sqlDB, err := sql.Open("pgx", server.DBUri)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	newRedisClient := redisClient
	if newRedisClient == nil {
		newRedisClient = server.GetRedisClient()
	}

	cache, err := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageRedis,
		RedisConfig:          cache.NewRedisConfigWithClient(newRedisClient),
		InvalidateWhenUpdate: true,   // when you create/update/delete objects, invalidate cache
		CacheTTL:             100000, // 100s
		CacheMaxItemCnt:      20,     // if length of objects retrieved one single time exceeds this number, then don't cache
	})

	if err != nil {
		fmt.Println("Error creating caching layer: ", err)
		return nil, err
	}

	err = gormDB.Use(cache)
	if err != nil {
		fmt.Println("Error using caching layer: ", err)
		return nil, err
	}

	return gormDB, nil
}

type JSONB map[string]interface{}

func (data JSONB) Value() (driver.Value, error) {
	return json.Marshal(data)
}

func (data *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var newData interface{}
	err := json.Unmarshal(source, &newData)
	if err != nil {
		return err
	}

	*data, ok = newData.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
