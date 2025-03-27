package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	http_server "url-shortener/internal/transport/http-server"
	"url-shortener/pkg/postgres"
)

type Config struct {
	Env            string             `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath    string             `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	PostgresConfig postgres.Config    `yaml:"postgres_config" env:"POSTGRES_CONFIG" env-required:"true"`
	RouterConfig   http_server.Config `yaml:"router_config" env:"ROUTER_CONFIG" env-required:"true"`
}

func NewConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file /{%s}/ does not exist", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", configPath)
	}

	return &config
}
