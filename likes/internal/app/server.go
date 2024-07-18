package app

import (
	"likes/config"
	grpcapp "likes/internal/app/grpc_app"
	"likes/internal/database/postgres"
	"log"
	"log/slog"
	"os"

	"gorm.io/gorm"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	DB         *gorm.DB
	Env        *config.Config
	GRPCServer *grpcapp.GRPC
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
	log := setupLogger(Env.Env)
	app.GRPCServer = grpcapp.NewGRPC(log, Env, PostgresClient)
	app.Env = Env
	app.DB = PostgresClient
	return *app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
