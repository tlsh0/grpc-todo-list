package main

import (
	"context"
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	proto "github.com/tlsh0/grpc-todo-list/task-service/proto"
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
