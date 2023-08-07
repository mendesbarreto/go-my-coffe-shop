package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	configs "github.com/mendesbarreto/go-my-coffe-shop/pkg/config"
)

type Config struct {
	configs.App     `yaml:"app"`
	configs.Http    `yaml:"http"`
	configs.MongoDb `yaml:mongodb`
}

var configInstance *Config

func GetConfig() *Config {

	dir, err := os.Getwd()

	if err != nil {
		panic("The os Getwd command is not support")
	}

	if configInstance == nil {
		fmt.Printf("Dir: %v", dir)
		if err := cleanenv.ReadConfig(dir+"config.yaml", &configInstance); err != nil {
			panic(err)
		}

	}

	return configInstance
}
