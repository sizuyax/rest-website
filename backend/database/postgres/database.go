package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"simple/backend/config"
)

var DB *sqlx.DB

func InitDB(cfg config.Config) error {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresName,
		"disable")

	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		logrus.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Successfully connected to the database!\n\n")

	return nil
}
