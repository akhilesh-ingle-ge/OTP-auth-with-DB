package config

import (
	"fmt"
	"jwt-auth/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "postgres://postgres:12345@localhost:5432/otpverify?sslmode=disable"

func SetUpDB() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{})
	DB = db
	fmt.Println("Connected to database")
}
