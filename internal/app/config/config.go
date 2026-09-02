package config

import (
	"log"

	"github.com/Hurricane199/catalog-service/internal/app/config/section"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Repository section.Repository
	Processor  section.Processor
	Monitor    section.Monitor
}

var Root Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Print(err.Error())
	}

	if err := envconfig.Process("APP", &Root); err != nil {
		log.Fatal(err.Error())
	}
}
