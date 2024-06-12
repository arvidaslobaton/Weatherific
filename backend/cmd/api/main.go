package main

import (
	"log"
	"os"

	"github.com/arvidaslobaton/Weatherific/backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}
	// port := "8081"
	server := server.NewServer(port)
	server.Start()
}