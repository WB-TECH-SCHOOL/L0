package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"l0/internal/cache"
	"l0/internal/delivery/docs"
	"l0/internal/delivery/routers"
	"l0/internal/infrastructure/config"
	"l0/internal/infrastructure/database"
	"l0/pkg/log"
)

func main() {
	router := gin.Default()

	logger, loggerFile := log.InitLoggers()
	defer loggerFile.Close()
	logger.Info().Msg("Logger Initialized")

	config.InitConfig()
	logger.Info().Msg("Config Initialized")

	db := database.GetDB()
	logger.Info().Msg("Database Initialized")

	orderCache := cache.InitCache()
	orderCache.LoadData(db)
	logger.Info().Msg("Cache Initialized")

	routers.InitRouting(router, db, orderCache, logger)
	logger.Info().Msg("Routing Initialized")

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info().Msg("Swagger Initialized")

	if err := router.Run("0.0.0.0:80"); err != nil {
		panic(fmt.Sprintf("Failed to run client: %s", err.Error()))
	}
}
