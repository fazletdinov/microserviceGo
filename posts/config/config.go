package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string      `yaml:"env" env:"ENV"`
	PostsServer PostsServer `yaml:"posts_server"`
	PostgresDB  PostgresDB  `yaml:"postgres_posts_db"`
	RedisDB     RedisDB     `yaml:"redis_posts_db"`
	Clients     GRPCClient  `yaml:"clients"`
	GRPC        GRPC        `yaml:"grpc"`
}

type PostsServer struct {
	PostsPort string `yaml:"posts_port" env:"POSTS_PORT"`
}

type PostgresDB struct {
	User     string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     uint   `yaml:"port" env:"POSTGRES_PORT"`
	Name     string `yaml:"name" env:"POSTGRES_NAME"`
	SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_USE_SSL"`
}

type RedisDB struct {
	Host string `yaml:"host" env:"REDIS_HOST"`
	Port uint   `yaml:"port" env:"REDIS_PORT"`
	Exp  uint   `yaml:"exp" env:"REDIS_EXPIRATION"`
}

type GRPCClient struct {
	Address string `yaml:"address"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func InitConfig() (*Config, error) {
	var env Config
	errEnv := godotenv.Load()
	if errEnv != nil {
		return nil, fmt.Errorf("ошибка при загрузки ENV %v", errEnv)
	}
	path := parseCommand()
	err := cleanenv.ReadConfig(path, &env)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении config.yaml %v", err)
	}
	return &env, nil
}

func parseCommand() string {
	var cfgPath string
	fset := flag.NewFlagSet("Notes", flag.ContinueOnError)
	fset.StringVar(&cfgPath, "path", os.Getenv("PATH_CONFIG"), "path to config file")
	fset.Parse(os.Args[1:])
	return cfgPath
}
