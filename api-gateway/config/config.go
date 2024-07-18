package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env               string            `yaml:"env" env:"ENV"`
	GatewayGRPCServer GatewayGRPCServer `yaml:"gatewayGRPCServer"`
	JWTConfig         JWTConfig         `yaml:"jwt_config"`
}

type GatewayGRPCServer struct {
	ApiGatewayPort     string `yaml:"api_gateway_port"`
	AuthServerAddress  string `yaml:"auth_server_address"`
	PostsServerAddress string `yaml:"posts_server_address"`
	LikesServerAddress string `yaml:"likes_server_address"`
}

type JWTConfig struct {
	PathPrivateKey    string `yaml:"path_private_key" env:"PATH_PRIVATE_KEY"`
	PathPublicKey     string `yaml:"path_public_key" env:"PATH_PUBLIC_KEY"`
	AccessTokenExp    uint   `yaml:"access_token_exp" env:"ACCESS_TOKEN_EXP"`
	RefreshTokenExp   uint   `yaml:"refresh_token_exp" env:"REFRESH_TOKEN_EXP"`
	SessionCookieName string `yaml:"session_cookie_name" env:"SESSION_COOKIE_NAME"`
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
	flag.StringVar(&cfgPath, "path", "", "path to config file")
	flag.Parse()
	if cfgPath == "" {
		cfgPath = os.Getenv("PATH_CONFIG")
	}
	return cfgPath
}
