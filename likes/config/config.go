package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string      `yaml:"env" env:"ENV"`
	LikesServer LikesServer `yaml:"likes_server"`
	PostgresDB  PostgresDB  `yaml:"postgres_likes_db"`
	GRPC        GRPC        `yaml:"grpc"`
	Jaeger      Jaeger      `yaml:"jaeger"`
}

type LikesServer struct {
	LikesPort string `yaml:"likes_port" env:"LIKES_PORT"`
}

type PostgresDB struct {
	User     string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     uint   `yaml:"port" env:"POSTGRES_PORT"`
	Name     string `yaml:"name" env:"POSTGRES_DB"`
	SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_USE_SSL"`
}

type GRPC struct {
	LikesGRPCPort int           `yaml:"likes_grpc_port" env:"LIKES_GRPC_PORT"`
	Timeout       time.Duration `yaml:"timeout"`
}

type Jaeger struct {
	CollectorUrl string `yaml:"collector_url"`
	Application  string `yaml:"application"`
}

func InitConfig() (*Config, error) {
	var env Config
	path := parseCommand()
	err := cleanenv.ReadConfig(path, &env)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении config.yaml %v", err)
	}
	return &env, nil
}

func parseCommand() string {
	var cfgPath string
	flag.StringVar(&cfgPath, "path", "", "path to config file")
	flag.Parse()
	if cfgPath == "" {
		cfgPath = os.Getenv("PATH_CONFIG")
	}
	return cfgPath
}
