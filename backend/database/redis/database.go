package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"simple/backend/config"
)

var Client *redis.Client

func InitRedis(cfg config.Config) error {

	adr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPorts)
	Client = redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0,
	})

	if _, err := Client.Ping().Result(); err != nil {
		logrus.Error("error to connect redis", err)
		return err
	}

	logrus.Info("Successfully connected to redis.")

	return nil
}
