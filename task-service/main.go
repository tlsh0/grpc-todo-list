package main

import (
	"log"
	"net"

	"github.com/tlsh0/grpc-todo-list/task-service/proto"
	"google.golang.org/grpc"
)

func main() {
	InitDB()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaskServiceServer(grpcServer, &TaskServiceServer{})

	log.Println("task-service is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
