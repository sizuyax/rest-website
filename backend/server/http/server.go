package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"simple/backend/cmd/cleaner"
	"simple/backend/config"
	"simple/backend/database/postgres"
	"simple/backend/database/redis"
	"simple/backend/server/http/handlers"
)

func InitWebServer() error {
	e := echo.New()

	cfg, err := config.EnvLoader()
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err = redis.InitRedis(cfg); err != nil {
		logrus.Error(err)
		return err
	}

	if err = postgres.InitDB(cfg); err != nil {
		logrus.Error(err)
		return err
	}

	if err = postgres.Migrate(postgres.DB, cfg); err != nil {
		logrus.Error(err)
		return err
	}

	go func() {
		if err = cleaner.Cleaner(postgres.DB); err != nil {
			logrus.Fatal(err)
		}
	}()

	handlers.Routes(e)

	return e.Start(":1323")
}
