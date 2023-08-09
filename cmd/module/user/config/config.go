package config

import (
	"github.com/caarlos0/env/v9"
)

type (
	Module struct {
		Name    string `env:"APP_NAME" envDefault:"module-user"`
		Version string `env:"APP_VERSION" envDefault:"0.0.1"`
	}

	Grcp struct {
		Host string `env-required:"true" env:"HTTP_HOST" envDefault:"localhost"`
		Port string `env-required:"true" env:"HTTP_PORT" envDefault:"3000"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL" envDefault:"debug"`
	}

	MongoDb struct {
		URL string `env-required:"true" env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
	}

	Config struct {
		Module
		Http
		Log
		MongoDb
	}
)

func GetConfig() *Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return &config
}
