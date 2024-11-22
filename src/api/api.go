package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfaxmodem/web-api/api/routers"
	"github.com/mfaxmodem/web-api/config"
	"github.com/mfaxmodem/web-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	RegisterSwagger(r, cfg)

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

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Golang Web API"
	docs.SwaggerInfo.Description = "Golang API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
