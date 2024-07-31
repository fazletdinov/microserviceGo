package app

import (
	"auth/config"
	grpcapp "auth/internal/app/grpc"
	"auth/internal/database/postgres"
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

func NewApp() Application {

	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	Client, errPostgres := postgres.NewPostgresClient(Env)
	if errPostgres != nil {
		log.Fatalf("ошибка подключения к Postgres - %v", errPostgres)
	}
	log := setupLogger(Env.Env)
	app.GRPCServer = grpcapp.NewGRPC(log, Env, Client)
	app.Env = Env
	app.DB = Client
	return *app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
		log = slog.New(slog.NewTextHandler(os.Stdout, opts))
	case envDev:
		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
		log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	case envProd:
		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}
		log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}
	slog.SetDefault(log)
	return log
}
