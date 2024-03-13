package cleaner

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

func Cleaner(db *sqlx.DB) error {

	logrus.Info("Cleaning up the database...")

	//threshold := time.Now().AddDate(0, 0, -7)
	threshold := time.Now().Add(-2 * time.Minute)

	if _, err := db.Exec(`DELETE FROM todo_list WHERE created_at < $1`, threshold); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Database cleaned up!\n\n")

	time.Sleep(2 * time.Minute)

	return nil
}
