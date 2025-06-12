package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tlsh0/grpc-todo-list/apigateway-service/proto"
	taskpb "github.com/tlsh0/grpc-todo-list/task-service/proto"
	userpb "github.com/tlsh0/grpc-todo-list/user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	userConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	taskConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	server := &ApiGatewayServer{
		UserClient: userpb.NewUserServiceClient(userConn),
		TaskClient: taskpb.NewTaskServiceClient(taskConn),
	}

	grpcServer := grpc.NewServer()
	proto.RegisterApiGatewayServiceServer(grpcServer, server)
	// Example: Register the task service as well, if needed
	// taskpb.RegisterTaskServiceServer(grpcServer, server) // Uncomment and implement if ApiGatewayServer implements TaskServiceServer

	go func() {
		lis, _ := net.Listen("tcp", ":50050")
		log.Fatal(grpcServer.Serve(lis))
	}()

	mux := runtime.NewServeMux()
	proto.RegisterApiGatewayServiceHandlerServer(context.Background(), mux, server)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
