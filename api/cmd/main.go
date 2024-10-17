package main

import (
	"c-ademy/api"
	"c-ademy/internal/config"
	"c-ademy/internal/db/schema"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var serverConfig config.Environment

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("cannot load the environment variables from file: %v", err)
		}
	}

	// Get the configuration
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("cannot get the configuration: %v", err)
	}

	serverConfig = *config

	dbPath := serverConfig.DbPath

	// Create the database file
	schemaManager, err := schema.NewSchemaManager(dbPath)
	if err != nil {
		log.Fatalf("cannot create schema manager: %v", err)
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		// Populate the database with some data
		err = schemaManager.Seed()
		if err != nil {
			log.Fatalf("cannot seed the database: %v", err)
		}
	}
}

func main() {
	e := api.GetRouter(&serverConfig)

	e.Logger.Fatal(e.Start(":1323"))
}
