package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"simple/backend/config"
	_ "simple/backend/database/postgres/migrations"
)

func Migrate(db *sqlx.DB, cfg config.Config) error {

	if cfg.Reload {

		logrus.Info("Reloading database...")

		if err := goose.DownTo(db.DB, ".", 0); err != nil {
			logrus.Fatal(err)
			return err
		}
	}

	if err := goose.Up(db.DB, "."); err != nil {
		logrus.Fatal(err)
		return err
	}

	return nil
}
