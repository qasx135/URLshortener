package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
	"url-shortener/pkg/postgres"
)

type Config struct {
	Env            string          `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath    string          `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer     HTTPServer      `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
	PostgresConfig postgres.Config `yaml:"postgres_config" env:"POSTGRES_CONFIG" env-required:"true"`
}

type HTTPServer struct {
	Addr        string        `yaml:"address" env:"HTTP_ADDR" env-default:"localhost:8080" env-required:"true"`
	TimeOut     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
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
