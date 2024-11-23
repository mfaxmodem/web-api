package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfaxmodem/web-api/docs"
	"github.com/mfaxmodem/web-api/src/api/routers"
	"github.com/mfaxmodem/web-api/src/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	// تنظیم Swagger
	ConfigureSwagger(cfg)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// مسیرهای نسخه‌بندی‌شده
	v1 := r.Group("api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	// شروع سرور
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

func ConfigureSwagger(cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
}
