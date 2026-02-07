package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AnuragDahiwade/task-manager/config"

	"github.com/AnuragDahiwade/task-manager/internal/user"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	log.Println("PostgreSQL connected")

	// Enable UUID extension
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`)

	// Auto migrate
	DB.AutoMigrate(&user.User{})

}
