package storage

import (
	. "circlesServer/modules/reader"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Redis() (*redis.Client, error) {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:" + GetConfig(`redis.PORT`)
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
