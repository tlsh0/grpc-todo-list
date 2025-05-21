package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
	ID          uint `gorm:"primaryKey`
	Title       string
	Description string
	Completed   bool
	Username    string //comes from decoded JWT
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=todo_app port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&Task{})
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	log.Println("Connected to Postgers and migrated task table")
}
