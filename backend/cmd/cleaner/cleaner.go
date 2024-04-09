package cleaner

import (
	"simple/backend/database/postgres"
	"simple/backend/logger"
	"time"
)

func CleanPostgres() error {

	logger.Logger.Debug("cleaning database...")

	//threshold := time.Now().AddDate(0, 0, -7)
	threshold := time.Now().Add(-2 * time.Minute)

	if _, err := postgres.DB.Exec(`DELETE FROM todo_list WHERE created_at < $1`, threshold); err != nil {
		logger.Logger.Error(err)
		return err
	}

	logger.Logger.Debug("database cleaned up!\n\n")

	time.Sleep(2 * time.Minute)

	return nil
}
