package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Config struct {
	MongoConfig  MongoConfig
	AppConstants AppConstants
}

func NewConfig() *Config {
	return &Config{}
}

type MongoConfig struct {
	ClusterName string `env:"MONGO_CLUSTER,required"`
	UserName    string `env:"MONGO_USERNAME,required"`
	Password    string `env:"MONGO_PASSWORD,required"`
	DBName      string `env:"MONGO_DATABASE_NAME,required"`
}

type AppConstants struct {
	PostCollectionName    string
	CommentCollectionName string
}

func (cfg *Config) loadEnv() *Config {
	err := godotenv.Load("./configs/.env")

	if err != nil {
		slog.Error(err.Error())
	}

	err = env.Parse(cfg)

	if err != nil {
		slog.Error(err.Error())
	}
	return cfg
}

func (cfg *Config) initAppConstants() {
	cfg.AppConstants.PostCollectionName = "Posts"
	cfg.AppConstants.CommentCollectionName = "Comments"
}

func (cfg *Config) InitConfig() *Config {
	cfg.loadEnv()
	cfg.initAppConstants()
	return cfg
}
