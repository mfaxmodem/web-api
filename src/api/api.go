package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfaxmodem/web-api/api/routers"
	"github.com/mfaxmodem/web-api/config"
)

func InitServer() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	r := gin.New()
	//RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery())

	// Add other routes here...
	RegisterRoutes(r)

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}
