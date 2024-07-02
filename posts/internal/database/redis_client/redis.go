package redis_client

import (
	"fmt"
	"posts/config"

	"github.com/go-redis/redis"
)

func InitRedisDB(env *config.Config) (*redis.Client, error) {
	redisURI := fmt.Sprintf("%s:%d", env.RedisDB.Host, env.RedisDB.Port)

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: "",
		DB:       0,
	})

	status := RedisClient.Ping()
	if status.Val() == "PONG" {
		return RedisClient, nil
	} else {
		return nil, fmt.Errorf("ошибка при подключении к Redis - %v", status)
	}
}
