package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tlsh0/grpc-todo-list/user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	go startGRPCServer()
	startHTTPGateway()
}

func startHTTPGateway() {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := proto.RegisterUserServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051", // This is where your gRPC server is running
		opts,
	)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	log.Println("HTTP Gateway listening on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Failed to serve HTTP gateway: %v", err)
	}
}

func startGRPCServer() {
	InitDB()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("task-service is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
