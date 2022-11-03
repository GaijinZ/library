package redisdb

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v9"
)

func RedisSetup() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := net.Dial("tcp", "192.168.1.55:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				conn.Close()
			}

			return conn, err
		},
		MaxIdleConns:    80,
		ConnMaxLifetime: 12000,
		PoolSize:        100,
	})

	checkRedisConnection(rdb)

	return rdb
}

func checkRedisConnection(rdb *redis.Client) {
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("error: %s", err)
	}

	log.Print(pong)
}
