package main

import (
	"log"
	"os"

	"coverage-service/backend/database"
	"coverage-service/backend/models"
	"coverage-service/backend/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.ConnectDB()
	database.DB.AutoMigrate(&models.Coverage{})

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Coverage service running on port:", port)
	r.Run(":" + port)
}
