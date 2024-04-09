package logger

import (
	"github.com/sirupsen/logrus"
	"simple/backend/config"
)

var Logger *logrus.Logger

func NewLogger(cfg config.Config) error {
	l := logrus.New()

	parsedLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Error("failed to parse log level, log level will be set [info]")
		return err
	}

	l.SetLevel(parsedLevel)

	Logger = l

	return nil
}
