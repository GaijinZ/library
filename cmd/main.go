package main

import (
	"github.com/GaijinZ/grpc/pkg/grpc/server"
	"github.com/GaijinZ/grpc/pkg/postgres"
	"github.com/GaijinZ/grpc/pkg/redisdb"
)

func main() {
	redisdb.RedisSetup()
	postgres.SetUpPostgres()
	server.SetupGRPCServer("8080")
}
