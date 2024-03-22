package main

import (
	"os"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/config"
)

func init() {
}

func main() {
	server := config.NewServer(os.Getenv("PORT"), nil)
	server.Run()
}
