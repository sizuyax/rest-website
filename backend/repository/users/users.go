package users

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func AddUser(db *sqlx.DB, username string, hashedPassword string) error {
	logrus.Info("Adding new user...")

	stmt, err := db.Prepare("INSERT INTO users (username, hashed_password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(username, hashedPassword); err != nil {
		return err
	}

	logrus.Info("New user added!\n\n")

	return nil
}

func GetHashedPassword(db *sqlx.DB, username string) (string, error) {

	var hashedPassword []byte

	if err := db.QueryRow("SELECT hashed_password FROM users WHERE username = $1", username).Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return string(hashedPassword), nil
}

func IsUserExists(db *sqlx.DB, username string) (bool, error) {

	smtm, err := db.Prepare("SELECT COUNT(*) FROM users WHERE username=$1")
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	defer smtm.Close()

	var count int
	if err = smtm.QueryRow(username).Scan(&count); err != nil {
		logrus.Error(err)
		return false, err
	}

	if count == 1 {
		logrus.Info("User already in db!")
		return true, err
	}

	return false, nil
}
