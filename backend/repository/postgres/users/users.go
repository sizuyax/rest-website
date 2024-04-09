package users

import (
	"database/sql"
	"simple/backend/database/postgres"
	"simple/backend/logger"
)

func AddUser(username string, hashedPassword string) error {

	stmt, err := postgres.DB.Prepare("INSERT INTO users (username, hashed_password) VALUES ($1, $2)")
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(username, hashedPassword); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}

func GetHashedPassword(username string) (string, error) {

	var hashedPassword []byte

	if err := postgres.DB.QueryRow("SELECT hashed_password FROM users WHERE username = $1", username).Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		logger.Logger.Error(err)

		return "", err
	}

	return string(hashedPassword), nil
}

func IsUserExists(username string) (bool, error) {

	smtm, err := postgres.DB.Prepare("SELECT COUNT(*) FROM users WHERE username=$1")
	if err != nil {
		logger.Logger.Error(err)
		return false, err
	}
	defer smtm.Close()

	var count int
	if err = smtm.QueryRow(username).Scan(&count); err != nil {
		logger.Logger.Error(err)
		return false, err
	}

	if count == 1 {
		return true, err
	}

	return false, nil
}
