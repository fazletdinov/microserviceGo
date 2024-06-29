package database

import (
	"fmt"
	"posts/config"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedisDB() error {
	redisURI := fmt.Sprintf("%s:%d", config.ConfigEnvs.RedisDB.Host, config.ConfigEnvs.RedisDB.Port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: "",
		DB:       0,
	})

	status := RedisClient.Ping()
	if status.Val() == "PONG" {
		return nil
	} else {
		return fmt.Errorf("ошибка при подключении к Redis - %v", status)
	}
}
