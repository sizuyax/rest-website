package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func EnvLoader() (Config, error) {

	logrus.Info("Loading environment variables...")

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Error(err)
		return Config{}, err
	}

	logrus.Info("Successfully loaded environment variables!\n\n")

	return cfg, nil
}
