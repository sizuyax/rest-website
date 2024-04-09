package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"admin"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresName     string `env:"POSTGRES_DB" envDefault:"database"`
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPorts    string `env:"POSTGRES_PORTS" envDefault:"5432"`
	RedisHost        string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPorts       string `env:"REDIS_PORTS" envDefault:"6379"`
	Reload           bool   `env:"RELOAD" envDefault:"false"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
	EchoPort         string `env:"ECHO_PORT" envDefault:":1323"`
}

func EnvLoader() (Config, error) {
	if err := godotenv.Load(); err != nil {
		logrus.Error("error loading .env file: ", err)
		return Config{}, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		logrus.Error(err)
		return Config{}, err
	}

	return *cfg, nil
}
