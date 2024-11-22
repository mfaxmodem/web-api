package main

import (
	"log"

	"github.com/mfaxmodem/web-api/api"
	"github.com/mfaxmodem/web-api/config"
	"github.com/mfaxmodem/web-api/data/cache"
)

func main() {
	// بارگذاری پیکربندی
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// اتصال به Redis
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	// راه‌اندازی سرور API
	api.InitServer()
}
