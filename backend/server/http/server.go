package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"simple/backend/cmd/cleaner"
	"simple/backend/config"
	"simple/backend/database/postgres"
	"simple/backend/database/redis"
	"simple/backend/logger"
	"simple/backend/server/http/routes"
)

type Server struct {
	e    *echo.Echo
	port string
}

func InitWebServer() (*Server, error) {
	e := echo.New()

	cfg, err := config.EnvLoader()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if err := logger.NewLogger(cfg); err != nil {
		logrus.Error(err)
		return nil, err
	}

	if err := postgres.InitPostgres(cfg); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	if err := postgres.MigratePostgres(cfg); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	if err := redis.InitRedis(cfg); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	go func() {
		if err := cleaner.CleanPostgres(); err != nil {
			logger.Logger.Warn(err)
		}
	}()

	routes.SetupRoutes(e)

	return &Server{
		e:    e,
		port: cfg.EchoPort,
	}, nil
}

func (s Server) StartServer() error {

	if err := s.e.Start(s.port); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}
