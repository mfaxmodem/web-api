package main

import (
	"log"

	_ "github.com/mfaxmodem/web-api/docs"
	"github.com/mfaxmodem/web-api/src/api"
	"github.com/mfaxmodem/web-api/src/config"
	"github.com/mfaxmodem/web-api/src/data/cache"
	"github.com/mfaxmodem/web-api/src/data/db"
	migration "github.com/mfaxmodem/web-api/src/data/db/migrations"

	"github.com/mfaxmodem/web-api/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	// Load configuration from config file or environment variables.
	cfg, err := config.GetConfig()
	logger := logging.NewLogger(cfg)

	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Redis Connection
	if err := cache.InitRedis(cfg); err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	// Database Connection
	if err := db.InitDb(cfg); err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()

	migration.Up1()

	// API Connection
	api.InitServer(cfg)
}
