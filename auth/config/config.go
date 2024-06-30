package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV"`
	AuthServer AuthServer `yaml:"auth_server"`
	MongoDB    MongoDB    `yaml:"mongo_auth_db"`
	PostgresDB PostgresDB `yaml:"postgres_auth_db"`
	JWTConfig  JWTConfig  `yaml:"jwt_config"`
}

type AuthServer struct {
	AuthPort string `yaml:"auth_port" env:"AUTH_PORT"`
}

type JWTConfig struct {
	PathPrivateKey    string `yaml:"path_private_key" env:"PATH_PRIVATE_KEY"`
	PathPublicKey     string `yaml:"path_public_key" env:"PATH_PUBLIC_KEY"`
	AccessTokenExp    uint   `yaml:"access_token_exp" env:"ACCESS_TOKEN_EXP"`
	RefreshTokenExp   uint   `yaml:"refresh_token_exp" env:"REFRESH_TOKEN_EXP"`
	SessionCookieName string `yaml:"session_cookie_name" env:"SESSION_COOKIE_NAME"`
}

type MongoDB struct {
	User     string `yaml:"user" env:"MONGO_INITDB_ROOT_USERNAME"`
	Password string `yaml:"password" env:"MONGO_INITDB_ROOT_PASSWORD"`
	Port     uint   `yaml:"port" env:"MONGO_PORT"`
	Host     string `yaml:"host" env:"MONGO_HOST"`
	CtxExp   uint   `yaml:"ctx_exp" env:"MONGO_CTX_EXP"`
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
	fset := flag.NewFlagSet("Auth", flag.ContinueOnError)
	fset.StringVar(&cfgPath, "path", os.Getenv("PATH_CONFIG"), "path to config file")
	fset.Parse(os.Args[1:])
	return cfgPath
}
