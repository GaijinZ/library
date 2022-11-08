package main

import (
	"github.com/GaijinZ/grpc/pkg/grpc/server"
	"github.com/GaijinZ/grpc/pkg/redisdb"
)

func main() {
	redisdb.RedisSetup()
	server.SetupGRPCServer("8080")
}
