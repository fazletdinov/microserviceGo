package grpcapp

import (
	"api-grpc-gateway/config"
	"log"
)

type Application struct {
	Env *config.Config
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	app.Env = Env
	return *app
}
