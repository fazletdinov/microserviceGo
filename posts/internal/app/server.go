package app

import (
	"github.com/go-redis/redis"
	"log"
	"posts/config"
	"posts/internal/database/postgres"
	"posts/internal/database/redis_client"

	"gorm.io/gorm"
)

type Application struct {
	DB    *gorm.DB
	Env   *config.Config
	Redis *redis.Client
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	PostgresClient, errPostgres := postgres.InitDatabse(Env)
	if errPostgres != nil {
		log.Fatalf("ошибка подключения к Postgres - %v", errPostgres)
	}
	RedisClient, errRedis := redis_client.InitRedisDB(Env)
	if errRedis != nil {
		log.Fatalf("ошибка подключения к Redis - %v", errPostgres)
	}
	app.Env = Env
	app.DB = PostgresClient
	app.Redis = RedisClient
	return *app
}
