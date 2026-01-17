package main

import (
	"context"
	"log"

	pb "github.com/JasperRosales/ombv1-common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpc *grpc.Server) {
	handler := &grpcHandler{}

	pb.RegisterOrderServiceServer(grpc, handler)

}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("[Info] New Order Received")
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
