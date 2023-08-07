package configs

type (
	App struct {
		Name    string `env-required:"true" env:"APP_NAME"`
		Version string `env-required:"true" env:"APP_VERSION"`
	}

	Http struct {
		Host string `env-required:"true" env:"HTTP_HOST"`
		Port string `env-required:"true" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL"`
	}

	MongoDb struct {
		URL string `env-required:"true" env"MONGO_URL"`
	}
)
