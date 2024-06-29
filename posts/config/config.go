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
	PostsServer PostsServer `yaml:"posts_server"`
	PostgresDB  PostgresDB  `yaml:"postgres_posts_db"`
	RedisDB     RedisDB     `yaml:"redis_posts_db"`
}

type PostsServer struct {
	PostsPort uint `yaml:"posts_port" env:"POSTS_PORT"`
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

var ConfigEnvs Config

func InitConfig() error {
	errEnv := godotenv.Load("")
	if errEnv != nil {
		return fmt.Errorf("ошибка при загрузки ENV %v", errEnv)
	}
	path := parseCommand()
	err := cleanenv.ReadConfig(path, &ConfigEnvs)
	if err != nil {
		return fmt.Errorf("ошибка при чтении config.yaml %v", err)
	}
	return nil
}

func parseCommand() string {
	var cfgPath string
	fset := flag.NewFlagSet("Notes", flag.ContinueOnError)
	fset.StringVar(&cfgPath, "path", os.Getenv("PATH_CONFIG"), "path to config file")
	fset.Parse(os.Args[1:])
	return cfgPath
}
