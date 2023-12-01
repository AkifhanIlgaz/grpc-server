package config

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectRedis(uri string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: uri,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	err := client.Set("test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")
	return client
}
