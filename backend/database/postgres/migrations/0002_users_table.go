package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func init() {
	goose.AddMigration(upUsersTable, downUsersTable)
}

func upUsersTable(tx *sql.Tx) error {

	logrus.Info("Creating users table...")

	if _, err := tx.Exec(`CREATE TABLE IF NOT EXISTS users (
    	id SERIAL,
    	username VARCHAR(255) NOT NULL PRIMARY KEY,
		hashed_password VARCHAR(255) NOT NULL
	)`); err != nil {
		logrus.Fatal(err)
		return err
	}

	logrus.Info("Users table created\n\n")

	return nil
}

func downUsersTable(tx *sql.Tx) error {

	logrus.Info("Dropping users table...")

	if _, err := tx.Exec(`DROP TABLE users`); err != nil {
		logrus.Fatal(err)
		return err
	}

	logrus.Info("Users table dropped\n\n")

	return nil
}
