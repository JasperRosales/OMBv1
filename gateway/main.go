package main

import (
	"log"
	"net/http"

	common "github.com/JasperRosales/ombv1-common"
	pb "github.com/JasperRosales/ombv1-common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr         = common.EnvString("ADDR", ":8080")
	orderService = common.EnvString("ORDER_SERVICE", "localhost:8081")
)

func main() {
	conn, err := grpc.NewClient(orderService, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("[Error] Failed to connect to order service: %v", err)
	}

	defer conn.Close()

	log.Println("[Info] Order Service is at ", orderService)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("[Info] Starting server at http://localhost:%s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("[Error] Failed to start the server!")
	}
}
