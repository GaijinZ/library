package main

import (
	"github.com/GaijinZ/grpc/grpc/server"
	"github.com/GaijinZ/grpc/redis"
)

func main() {
	redis.RedisSetup()
	server.SetupGRPCServer("8080")
}
