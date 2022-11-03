package main

import (
	"github.com/GaijinZ/grpc/grpc/server"
	"github.com/GaijinZ/grpc/redisdb"
)

func main() {
	redisdb.RedisSetup()
	server.SetupGRPCServer("8080")
}
