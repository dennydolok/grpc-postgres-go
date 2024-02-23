package database

import (
	"grpc-go2/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:bearuang16@localhost:5432/local"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("error : %v", err)
	}

	db.AutoMigrate(&models.Book{})
	return db
}
