package main

import (
	"log"
	"logistics-service/database"
	"logistics-service/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, gunakan environment variables")
	}

	database.Connect()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run(":" + port)
}
