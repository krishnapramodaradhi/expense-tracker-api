package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/config"
)

func init() {
}

func main() {
	log.Println(os.Getenv("PORT"), os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("There was an error while loading .env variables", err)
		}
	}
	server := config.NewServer(os.Getenv("PORT"), nil)
	server.Run()
}
