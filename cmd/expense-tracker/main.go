package main

import (
	"log"
	"os"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/config"
)

func init() {
}

func main() {
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("there was an error while connecting to the database", err)
	}
	server := config.NewServer(os.Getenv("PORT"), db.DB)
	server.Run()
}
