package main

import (
	"context"
	"time"

	proto "github.com/tlsh0/grpc-todo-list/user-service/proto"

	"github.com/golang-jwt/jwt/v5"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
}

var jwtKey = []byte("secret-key")

func (s *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	// TODO: save to DB here
	token := generateToken(req.Username)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	// TODO: validate user here
	token := generateToken(req.Username)
	return &proto.AuthResponse{Token: token}, nil
}

func generateToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(jwtKey)
	return t
}
