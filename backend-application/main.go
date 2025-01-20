package main

import (
	"flight-app/config"
	"flight-app/routers"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("No .env file found")
	}
	_, err := config.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}
	router := routers.SetupRouter()
	if err := router.Run(":8090"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}

}
