package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"simple/backend/config"
	"simple/backend/logger"
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
		logger.Logger.Error("error to connect redis", err)
		return err
	}

	logger.Logger.Debug("successfully connected to redis!\n\n")

	return nil
}
