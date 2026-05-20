package main

import (
	"log"
	"os"

	"inventory-tracker/backend/database"
	"inventory-tracker/backend/models"
	"inventory-tracker/backend/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Item{})
	// database.DB.AutoMigrate(&models.Coverage{})

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
