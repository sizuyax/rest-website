package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func init() {
	goose.AddMigration(upTodoListTable, downTodoListTable)
}

func upTodoListTable(tx *sql.Tx) error {

	logrus.Info("Creating todo_list table...")

	if _, err := tx.Exec(`CREATE TABLE IF NOT EXISTS todo_list (
		id SERIAL,
		title VARCHAR(255) NOT NULL,
    	username VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Todo_list table created\n\n")

	return nil
}

func downTodoListTable(tx *sql.Tx) error {

	logrus.Info("Dropping todo_list table...")

	if _, err := tx.Exec(`DROP TABLE todo_list`); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Todo_list table dropped\n\n")

	return nil
}
