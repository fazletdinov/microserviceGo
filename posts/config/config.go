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
	PostsServer PostsServer `yaml:"posts_server"`
	PostgresDB  PostgresDB  `yaml:"postgres_posts_db"`
	RedisDB     RedisDB     `yaml:"redis_posts_db"`
	Clients     GRPCClient  `yaml:"clients"`
	GRPC        GRPC        `yaml:"grpc"`
	Jaeger      Jaeger      `yaml:"jaeger"`
}

type PostsServer struct {
	PostsPort string `yaml:"posts_port" env:"POSTS_PORT"`
}

type PostgresDB struct {
	User     string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     uint   `yaml:"port" env:"POSTGRES_PORT"`
	Name     string `yaml:"name" env:"POSTGRES_DB"`
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
	PostsGRPCPort int           `yaml:"posts_grpc_port" env:"POSTS_GRPC_PORT"`
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
