package redisdb

import (
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func AddToRedisDB(uid string, books []byte) {
	err := RedisSetup().Set(ctx, uid, books, 0).Err()
	if err != nil {
		log.Fatalf("Can not set data to redis %v", err)
	}
}

func GetFromRedisDB(uid string) {
	val, err := RedisSetup().Get(ctx, uid).Result()
	if err != nil {
		log.Fatalf("Can not get data get from redis %v", err)
	}

	fmt.Println(uid, val)
}

func DeleteFromRedisDB(uid string) {
	RedisSetup().Del(ctx, uid)
}
