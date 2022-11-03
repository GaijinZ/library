package redisdb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func RedisSetup() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.55:6379",
		Password: "",
		DB:       0,
	})

	checkRedisConnection(rdb)

	return rdb
}

func checkRedisConnection(rdb *redis.Client) {
	pong, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pong)
}
