package main

import (
	"log"
	"net"

	"github.com/tlsh0/grpc-todo-list/user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	InitDB()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("user-service is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
