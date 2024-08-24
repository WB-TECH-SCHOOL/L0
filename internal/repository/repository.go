package repository

import (
	"context"
	"l0/internal/models"
)

type Orders interface {
	Create(ctx context.Context, order models.Order) error
	GetAll(ctx context.Context) ([]models.Order, error)
}
