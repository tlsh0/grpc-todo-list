package main

import (
	"context"

	taskpb "github.com/tlsh0/grpc-todo-list/task-service/proto"
	userpb "github.com/tlsh0/grpc-todo-list/user-service/proto"

	"github.com/tlsh0/grpc-todo-list/apigateway-service/proto"
)

type ApiGatewayServer struct {
	proto.UnimplementedApiGatewayServiceServer
	UserClient userpb.UserServiceClient
	TaskClient taskpb.TaskServiceClient
}

func (s *ApiGatewayServer) RegisterUser(ctx context.Context, req *userpb.RegisterRequest) (*userpb.AuthResponse, error) {
	return s.UserClient.Register(ctx, req)
}
func (s *ApiGatewayServer) LoginUser(ctx context.Context, req *userpb.LoginRequest) (*userpb.AuthResponse, error) {
	return s.UserClient.Login(ctx, req)
}
func (s *ApiGatewayServer) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.TaskResponse, error) {
	return s.TaskClient.CreateTask(ctx, req)
}
func (s *ApiGatewayServer) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.TaskListResponse, error) {
	return s.TaskClient.ListTasks(ctx, req)
}
func (s *ApiGatewayServer) CompleteTask(ctx context.Context, req *taskpb.CompleteTaskRequest) (*taskpb.TaskResponse, error) {
	return s.TaskClient.CompleteTask(ctx, req)
}
func (s *ApiGatewayServer) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	return s.TaskClient.DeleteTask(ctx, req)
}
