package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=todo_app port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("Connected to DB and migrated schema.")
}
