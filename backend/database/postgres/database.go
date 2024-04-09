package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"simple/backend/config"
	"simple/backend/logger"
)

var DB *sqlx.DB

func InitPostgres(cfg config.Config) error {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresName,
		"disable",
	)

	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		logger.Logger.Error(err)
		return err
	}

	logger.Logger.Debug("successfully connected to postgres!\n\n")

	return nil
}
