package main

import (
	"context"

	"github.com/tlsh0/grpc-todo-list/apigateway-service/proto"
	taskpb "github.com/tlsh0/grpc-todo-list/task-service/proto"
	userpb "github.com/tlsh0/grpc-todo-list/user-service/proto"
)

type ApiGatewayServer struct {
	proto.UnimplementedApiGatewayServiceServer
	UserClient userpb.UserServiceClient
	TaskClient taskpb.TaskServiceClient
}

func (s *ApiGatewayServer) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	userReq := &userpb.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
	return s.UserClient.Register(ctx, userReq)
}
func (s *ApiGatewayServer) LoginUser(context.Context, *proto.LoginRequest) (*proto.AuthResponse, error) {
	return nil, nil
}
func (s *ApiGatewayServer) CreateTask(context.Context, *proto.CreateTaskRequest) (*proto.TaskResponse, error) {
	return nil, nil
}
func (s *ApiGatewayServer) ListTasks(context.Context, *proto.ListTasksRequest) (*proto.TaskListResponse, error) {
	return nil, nil
}
func (s *ApiGatewayServer) CompleteTask(context.Context, *proto.CompleteTaskRequest) (*proto.TaskResponse, error) {
	return nil, nil
}
func (s *ApiGatewayServer) DeleteTask(context.Context, *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	return nil, nil
}
