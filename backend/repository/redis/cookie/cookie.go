package cookie

import (
	"simple/backend/database/redis"
	"simple/backend/logger"
	"time"
)

func SetCookieSession(sessionID string, duration time.Duration) error {

	expirationSeconds := duration.Hours()

	if err := redis.Client.Set(sessionID, "active", time.Duration(expirationSeconds)*time.Hour).Err(); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}

func DeleteCookieFromRedis(sessionID string) error {

	if err := redis.Client.Del(sessionID).Err(); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}
