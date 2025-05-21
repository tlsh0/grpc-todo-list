package main

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/tlsh0/grpc-todo-list/user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	regResp, err := client.Register(ctx, &proto.RegisterRequest{
		Username: "testuser",
		Password: "password123",
	})
	if err != nil {
		log.Fatalf("Register failed: %v", err)
	}
	fmt.Println("Register -> JWT Token:", regResp.Token)

	loginResp, err := client.Login(ctx, &proto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	fmt.Println("Login -> JWT Token:", loginResp.Token)
}
