package main

import (
	"log"
	"net"

	proto "github.com/tlsh0/grpc-todo-list/user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &UserServer{})

	log.Println("UserService running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
