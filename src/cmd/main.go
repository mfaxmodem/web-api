package main

import (
	"log"

	_ "github.com/mfaxmodem/web-api/docs"
	"github.com/mfaxmodem/web-api/src/api"
	"github.com/mfaxmodem/web-api/src/config"
	"github.com/mfaxmodem/web-api/src/data/cache"
	"github.com/mfaxmodem/web-api/src/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	// Load configuration from config file or environment variables.
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Redis Connection
	cache.InitRedis(cfg)
	if err := cache.InitRedis(cfg); err != nil {
		log.Fatalf("failed to initialize Redis %v", err)
	}
	defer cache.CloseRedis()

	//Database Connection
	if err := db.InitDb(cfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.CloseDb()

	// API Connection
	api.InitServer(cfg)
}
