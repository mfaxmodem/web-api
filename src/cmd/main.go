package main

import (
	"log"

	"github.com/mfaxmodem/web-api/api"
	"github.com/mfaxmodem/web-api/config"
	"github.com/mfaxmodem/web-api/data/cache"
	"github.com/mfaxmodem/web-api/data/db"
)

func main() {
	// Load configuration from config file or environment variables.
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Redis Connection
	cache.InitRedis(cfg)
	if err := db.InitDb(cfg); err != nil {
		log.Fatalf("failed to initialize Redis %v", err)
	}
	defer cache.CloseRedis()

	//Database Connection
	if err := db.InitDb(cfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.CloseDb()

	// API Connection
	api.InitServer()
}
