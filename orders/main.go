package main

import (
	"context"
	"log"
	"net"

	common "github.com/JasperRosales/ombv1-common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:8082")
)

func main() {
	grpc := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("[Error] Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpc)

	svc.CreateOrder(context.Background())

	log.Println("GRPC Server started at ", grpcAddr)

	if err := grpc.Serve(l); err != nil {
		log.Fatalf("[Error] Failed to serve: %v", err)
	}
}
