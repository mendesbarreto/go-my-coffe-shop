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
		Host string `env-required:"true" env:"GRCP_HOST" envDefault:"0.0.0.0"`
		Port string `env-required:"true" env:"GRCP_PORT" envDefault:"3000"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL" envDefault:"debug"`
	}

	MongoDb struct {
		URL string `env-required:"true" env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
	}

	Config struct {
		Module
		Grcp
		Log
		MongoDb
		EnableGRPCReflection bool `env:"ENABLE_GRPC_REFLECTION" envDefault:"true"`
	}
)

func GetConfig() *Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return &config
}
