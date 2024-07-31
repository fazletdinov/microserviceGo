package grpcapp

import (
	"api-grpc-gateway/config"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	Env *config.Config
	Log *slog.Logger
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	log := setupLogger(Env.Env)
	app.Env = Env
	app.Log = log
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
