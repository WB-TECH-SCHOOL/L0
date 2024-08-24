package cache

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"l0/internal/errs"
	"l0/internal/infrastructure/config"
	"l0/internal/models"
	"l0/internal/repository"
	"sync"
	"time"
)

type Cache interface {
	LoadData(db *sqlx.DB)

	Add(order models.Order)
	GetByID(ID string) (models.Order, error)
}

type cache struct {
	cache map[string]json.RawMessage
	mu    sync.RWMutex
}

func InitCache() Cache {
	return &cache{
		cache: make(map[string]json.RawMessage),
		mu:    sync.RWMutex{},
	}
}

func (c *cache) LoadData(db *sqlx.DB) {
	ordersRepo := repository.InitOrderRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt(config.DBResponseTime))*time.Second)
	defer cancel()

	orders, err := ordersRepo.GetAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, order := range orders {
		c.cache[order.ID] = order.Data
	}
}

func (c *cache) Add(order models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[order.ID] = order.Data
}

func (c *cache) GetByID(ID string) (models.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, ok := c.cache[ID]
	if !ok {
		return models.Order{}, errs.ErrNoOrder
	}
	return models.Order{
		ID:   ID,
		Data: order,
	}, nil
}
