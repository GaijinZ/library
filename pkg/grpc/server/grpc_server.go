package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/GaijinZ/grpc/pkg/grpc/handlers"
	"github.com/GaijinZ/grpc/pkg/protobuff"

	"google.golang.org/grpc"
)

// set the grpc server up at given port
func SetupGRPCServer(port string) {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protobuff.RegisterLibraryServer(s, &handlers.BookServer{})
	log.Printf("server listining on port :%v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Starting grpc server...")
}
