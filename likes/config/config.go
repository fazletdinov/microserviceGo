package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string      `yaml:"env" env:"ENV"`
	LikesServer LikesServer `yaml:"likes_server"`
	PostgresDB  PostgresDB  `yaml:"postgres_likes_db"`
}

type LikesServer struct {
	LikesPort string `yaml:"likes_port" env:"LIKES_PORT"`
}

type PostgresDB struct {
	User     string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     uint   `yaml:"port" env:"POSTGRES_PORT"`
	Name     string `yaml:"name" env:"POSTGRES_NAME"`
	SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_USE_SSL"`
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
	fset := flag.NewFlagSet("Likes", flag.ContinueOnError)
	fset.StringVar(&cfgPath, "path", os.Getenv("PATH_CONFIG"), "path to config file")
	fset.Parse(os.Args[1:])
	return cfgPath
}
