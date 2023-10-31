package config

import (
	"github.com/caarlos0/env/v9"
)

type (
	Module struct {
		Name        string `env:"APP_NAME" envDefault:"module-user-go-local"`
		Version     string `env:"APP_VERSION" envDefault:"0.0.1"`
		AuthSecrete string `env:"SECRETE_AUTH_TOKEN" envDefault:"TopScrect"`
	}

	Grcp struct {
		Host string `env-required:"true" env:"GRCP_HOST" envDefault:"0.0.0.0"`
		Port string `env-required:"true" env:"GRCP_PORT" envDefault:"3000"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL" envDefault:"debug"`
	}

	MongoDb struct {
		URI string `env-required:"true" env:"MONGO_URI" envDefault:"mongodb://192.168.2.32:27017/?directConnection=true"`
	}

	Redis struct {
		URI string `env-required:"true" env:"REDIS_URI" envDefault:"192.168.2.32:6379"`
	}

	Config struct {
		Module
		Grcp
		Log
		MongoDb
		Redis
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
