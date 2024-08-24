package services

import (
	"github.com/rs/zerolog"
	"l0/internal/cache"
	"l0/internal/models"
	"l0/internal/repository"
	"l0/pkg/log"
	"time"
)

type orderService struct {
	cache          cache.Cache
	orderRepo      repository.Orders
	dbResponseTime time.Duration
	logger         *zerolog.Logger
}

func InitOrderService(
	cache cache.Cache,
	orderRepo repository.Orders,
	dbResponseTime time.Duration,
	logger *zerolog.Logger,
) Orders {
	return &orderService{
		cache:          cache,
		orderRepo:      orderRepo,
		dbResponseTime: dbResponseTime,
		logger:         logger,
	}
}

func (o *orderService) GetByID(ID string) (models.Order, error) {
	order, err := o.cache.GetByID(ID)
	if err != nil {
		o.logger.Error().Msg(err.Error())
		return models.Order{}, err
	}
	o.logger.Info().Msg(log.Normalizer(log.GetObject, log.Order, ID))
	return order, nil
}
