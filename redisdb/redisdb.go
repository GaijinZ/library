package redisdb

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v9"
)

var RedisPool *redis.Client

type Redis interface {
	Set(uid string, books []byte) (bool, error)
	Get(uid string) (string, error)
	Delete(uid string) (bool, error)
}

type Books struct{}

var ctx = context.Background()

func RedisSetup() {
	RedisPool = redis.NewClient(&redis.Options{
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := net.Dial("tcp", "192.168.1.55:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				conn.Close()
			}

			return conn, err
		},
	})

	checkRedisConnection(RedisPool)
}

func checkRedisConnection(rdb *redis.Client) {
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("error: %s", err)
	}

	log.Print(pong)

}

func (b Books) Set(uid string, books []byte) (bool, error) {
	err := RedisPool.Set(ctx, uid, books, 0).Err()
	if err != nil {
		return false, fmt.Errorf("could not set data %v", err)
	}

	return true, nil
}

func (b Books) Get(uid string) (string, error) {
	val, err := RedisPool.Get(ctx, uid).Result()
	if err != nil {
		return "", fmt.Errorf("can not get data get from redis %v", err)
	}

	return val, nil
}

func (b Books) Delete(uid string) (bool, error) {
	if err := RedisPool.Del(ctx, uid).Err(); err != nil {
		return false, fmt.Errorf("record not found")
	}

	return true, nil
}
