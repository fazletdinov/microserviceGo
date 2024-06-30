package app

import (
	"auth/config"
	"auth/internal/database/postgres"
	"log"

	"gorm.io/gorm"
)

type Application struct {
	DB  *gorm.DB
	Env *config.Config
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	Client, errPostgres := postgres.NewPostgresClient(Env)
	if errPostgres != nil {
		log.Fatalf("ошибка подключения к Postgres - %v", errPostgres)
	}
	app.Env = Env
	app.DB = Client
	return *app
}
