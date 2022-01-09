package storage

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var client *redis.Client

func Redis() (*redis.Client, error) {
	//Initializing redis
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:" + viper.GetString(`redis.PORT`)
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
