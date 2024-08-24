package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"l0/internal/cache"
	"l0/internal/delivery/handlers"
	"l0/internal/infrastructure/config"
	"l0/internal/repository"
	"l0/internal/services"
	"time"
)

func InitRouting(engine *gin.Engine, db *sqlx.DB, cache cache.Cache, logger *zerolog.Logger) {
	dbResponseTime := time.Duration(viper.GetInt(config.DBResponseTime)) * time.Second

	// Инициализация репозиториев
	orderRepo := repository.InitOrderRepository(db)

	// Инициализация сервисов
	orderService := services.InitOrderService(cache, orderRepo, dbResponseTime, logger)

	// Инициализация хендлеров
	orderHandler := handlers.InitOrderHandler(orderService)

	// Инициализация группы маршрутов
	baseGroup := engine.Group("/api")

	InitOrderRouter(baseGroup, orderHandler)
}

func InitOrderRouter(group *gin.RouterGroup, handler handlers.Orders) {
	orderGroup := group.Group("/order")

	orderGroup.GET(":order_id", handler.GetByID)
}
