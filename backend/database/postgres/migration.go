package postgres

import (
	"github.com/pressly/goose"
	"simple/backend/config"
	_ "simple/backend/database/postgres/migrations"
	"simple/backend/logger"
)

func MigratePostgres(cfg config.Config) error {

	if cfg.Reload {

		logger.Logger.Debug("reloading database...")

		if err := goose.DownTo(DB.DB, ".", 0); err != nil {
			logger.Logger.Error(err)
			return err
		}
	}

	if err := goose.Up(DB.DB, "."); err != nil {
		logger.Logger.Fatal(err)
	}

	logger.Logger.Debug("migrations successfully happened\n\n")

	return nil
}
