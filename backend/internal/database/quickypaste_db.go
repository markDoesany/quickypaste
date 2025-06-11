package database

import (
	"fmt"
	"log"
	"os"

	"github.com/smokeyghost/QuickyPaste/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
