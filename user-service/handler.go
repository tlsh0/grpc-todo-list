package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tlsh0/grpc-todo-list/user-service/proto"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
}

var jwtKey = []byte("secret-key")

func (s *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	log.Println("Registering new user:", req.Username)

	var existing User
	if err := db.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("username already taken")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}

	user := User{Username: req.Username, Password: string(hashedPwd)}
	if err := db.Create(&user).Error; err != nil {
		log.Printf("DB error during user creation: %v", err)
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	token := generateToken(req.Username)
	log.Printf("user registered in %v\n", req.Username)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {

	var user User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	token := generateToken(req.Username)
	log.Printf("user logged in %v\n", req.Username)
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
