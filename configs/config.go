package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Config struct {
	MongoConfig MongoConfig
}

type MongoConfig struct {
	ClusterName string `env:"MONGO_CLUSTER,required"`
	UserName    string `env:"MONGO_USERNAME,required"`
	Password    string `env:"MONGO_PASSWORD,required"`
}

func LoadEnv() *Config {
	err := godotenv.Load("./configs/.env")

	if err != nil {
		slog.Error(err.Error())
	}

	config := &Config{}

	err = env.Parse(config)

	if err != nil {
		slog.Error(err.Error())
	}
	return config
}
