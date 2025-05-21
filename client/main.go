package main

import (
	"context"
	"log"
	"time"

	proto "github.com/tlsh0/grpc-todo-list/task-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}
	defer conn.Close()

	// client := protoUser.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// register request
	// regResp, err := client.Register(ctx, &proto.RegisterRequest{
	// 	Username: "testuser3",
	// 	Password: "password123",
	// })
	// if err != nil {
	// 	log.Fatalf("Register failed: %v", err)
	// }
	// fmt.Println("Register -> JWT Token:", regResp.Token)

	//login request
	// loginResp, err := client.Login(ctx, &proto.LoginRequest{
	// 	Username: "testuser3",
	// 	Password: "password123",
	// })
	// if err != nil {
	// 	log.Fatalf("Login failed: %v", err)
	// }
	// fmt.Println("Login -> JWT Token:", loginResp.Token)

	task := proto.NewTaskServiceClient(conn)

	// token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc5MDU5OTIsInVzZXJuYW1lIjoidGVzdHVzZXIzIn0.LiO5V8kzkyGuI9qOACADWj1pz8UNXbRb7NFCTQe_Mqo`
}
