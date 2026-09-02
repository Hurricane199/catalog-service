package main

import (
	"log"

	"github.com/Hurricane199/catalog-service/internal/app/config"
)

func main() {
	config.Load()

	cfg := config.Root

	log.Printf("Server will start on port: %d", cfg.Processor.WebServer.ListenPort)
	log.Printf("Database: %s@%s/%s",
		cfg.Repository.Postgres.Username,
		cfg.Repository.Postgres.Address,
		cfg.Repository.Postgres.Name)
	log.Printf("Enviroment: %s, LogLevel: %s",
		cfg.Monitor.Enviroment,
		cfg.Monitor.LogLevel)
}
