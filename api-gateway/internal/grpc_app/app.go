package grpcapp

import (
	"api-grpc-gateway/config"
	"log"
	"os"

	"github.com/rs/zerolog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	Env *config.Config
	Log *zerolog.Logger
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

func setupLogger(env string) *zerolog.Logger {
	zerolog.TimeFieldFormat = "02/Jan/2006 - 15:04:05 -0700"
	switch env {
	case envLocal:
		logger := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "02/Jan/2006 - 15:04:05 -0700",
		}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envDev:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envProd:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	}
	return nil
}
