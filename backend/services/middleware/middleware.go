package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"net/http"
	"simple/backend/database/redis"
	"simple/backend/logger"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sessionID, err := c.Cookie("session")

		if err != nil || sessionID.Value == "" {
			return c.Redirect(http.StatusMovedPermanently, "/")
		}

		if _, err := redis.Client.Get(sessionID.Value).Result(); err != nil {
			return c.Redirect(http.StatusMovedPermanently, "/")
		}

		return next(c)
	}
}

func GenerateSessionID() (string, error) {
	token := make([]byte, 32)

	if _, err := rand.Read(token); err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	return base64.URLEncoding.EncodeToString(token), nil
}
