package main

import (
	"log"
	"os"
	"time-tracker/internal/handler"
	"time-tracker/internal/repository"
	"time-tracker/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading env variables.")
	}

	db, err := repository.ConnectDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})

	if err != nil {
		log.Fatalf("Failed to init DB: %s", err.Error())
	}

	repository.AutoMigrate(db)

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes().Run(":8080")
}
