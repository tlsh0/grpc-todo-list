package main

import (
	"context"
	"errors"
	"log"

	"github.com/tlsh0/grpc-todo-list/task-service/proto"

	"github.com/golang-jwt/jwt/v5"
)

type TaskServiceServer struct {
	proto.UnimplementedTaskServiceServer
}

var jwtKey = []byte("secret-key")

func parseJWT(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	username, ok := (*claims)["username"].(string)
	if !ok {
		return "", errors.New("username not found in token")
	}

	return username, nil
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskResponse, error) {
	username, err := parseJWT(req.Token)
	if err != nil {
		log.Println("JWT parse error:", err)
		return nil, err
	}

	task := Task{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		Username:    username,
	}

	result := DB.Create(&task)
	if result.Error != nil {
		log.Println("DB create error:", result.Error)
		return nil, result.Error
	}

	log.Printf("Task create for user %s: %s\n", username, task.Title)

	return &proto.TaskResponse{
		Task: &proto.Task{
			Id:          int32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		},
	}, nil
}

func (s *TaskServiceServer) ListTasks(ctx context.Context, req *proto.ListTasksRequest) (*proto.TaskListResponse, error) {
	username, err := parseJWT(req.Token)
	if err != nil {
		log.Println("JWT parse error:", err)
		return nil, err
	}

	var tasks []Task
	result := DB.Where("username = ?", username).Find(&tasks)
	if result.Error != nil {
		log.Println("DB query error:", result.Error)
		return nil, result.Error
	}

	var protoTasks []*proto.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, &proto.Task{
			Id:          int32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		})
	}

	return &proto.TaskListResponse{Tasks: protoTasks}, nil
}

func (s *TaskServiceServer) CompleteTask(ctx context.Context, req *proto.CompleteTaskRequest) (*proto.TaskResponse, error) {
	username, err := parseJWT(req.Token)
	if err != nil {
		log.Println("JWT parse error:", err)
		return nil, err
	}

	var task Task
	result := DB.First(&task, req.Id)
	if result.Error != nil {
		log.Println("Task not found:", result.Error)
		return nil, result.Error
	}

	if task.Username != username {
		log.Println("Unauthorized attempt to complete task")
		return nil, errors.New("unauthorized")
	}

	task.Completed = true
	DB.Save(&task)

	return &proto.TaskResponse{
		Task: &proto.Task{
			Id:          int32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		},
	}, nil
}

func (s *TaskServiceServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	username, err := parseJWT(req.Token)
	if err != nil {
		log.Println("JWT parse error:", err)
		return nil, err
	}

	var task Task
	result := DB.First(&task, req.Id)
	if result.Error != nil {
		log.Println("Task not found:", result.Error)
		return nil, result.Error
	}

	if task.Username != username {
		log.Println("Unauthorized attempt to complete task")
		return nil, errors.New("unauthorized")
	}

	if err := DB.Delete(&task).Error; err != nil {
		log.Println("Delete error:", err)
		return nil, err
	}

	log.Printf("Deleted task ID %d for user %s\n", task.ID, username)

	return &proto.DeleteTaskResponse{
		Message: "Task deleted succesfully",
	}, nil
}
