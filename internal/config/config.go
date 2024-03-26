package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Mongo Mongo
		HTTP  HTTP
		JWT   JWT
	}

	Mongo struct {
		URI        string `envconfig:"MONGO_URI"`
		Name       string `envconfig:"MONGO_DBNAME"`
		Collection string `envconfig:"MONGO_COLL"`
	}

	HTTP struct {
		Port string `envconfig:"HTTP_PORT"`
	}

	JWT struct {
		SecretKey string `envconfig:"JWT_SECRETKEY"`
	}
)

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
