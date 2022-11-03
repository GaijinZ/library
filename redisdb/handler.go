package redisdb

import (
	"context"
	"fmt"
	"log"
)

type Redis interface {
	Set(uid string, books []byte)
	Get(uid string)
	Delete(uid string)
}

type Books struct{}

var ctx = context.Background()

func (b Books) Set(uid string, books []byte) {
	err := RedisSetup().Set(ctx, uid, books, 0).Err()
	if err != nil {
		log.Fatalf("Can not set data to redis %v", err)
	}
}

func (b Books) Get(uid string) {
	val, err := RedisSetup().Get(ctx, uid).Result()
	if err != nil {
		log.Fatalf("Can not get data get from redis %v", err)
	}

	fmt.Println(uid, val)
}

func (b Books) Delete(uid string) {
	RedisSetup().Del(ctx, uid)
}
