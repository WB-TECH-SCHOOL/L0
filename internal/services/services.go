package services

import "l0/internal/models"

type Orders interface {
	GetByID(ID string) (models.Order, error)
}
