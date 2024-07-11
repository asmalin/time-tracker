package main

// @title Time Tracker API
// @version 1.0
// @description API для отслеживания времени выполнения задач

import (
	"embed"
	"os"
	"time-tracker/internal/handler"
	"time-tracker/internal/migrator"
	"time-tracker/internal/repository"
	"time-tracker/internal/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "time-tracker/docs"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

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

	log.Info("Successfully connected to database")

	migrator := migrator.MustGetNewMigrator(MigrationsFS, migrationsDir)

	err = migrator.ApplyMigrations(db)

	if err != nil {
		panic(err)
	}

	log.Info("Migration was completed successfully")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	log.Info("Starting server")
	handlers.InitRoutes(log).Run("0.0.0.0:" + os.Getenv("SERVER_PORT"))
}
