package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"l0/internal/models"
	"l0/pkg/customerr"
)

type orderRepository struct {
	db *sqlx.DB
}

func InitOrderRepository(db *sqlx.DB) Orders {
	return orderRepository{
		db: db,
	}
}

func (o orderRepository) Create(ctx context.Context, order models.Order) error {
	query := `INSERT INTO orders(id, data) VALUES ($1, $2)`

	res, err := o.db.ExecContext(ctx, query, order.ID, order.Data)
	if err != nil {
		return customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.ScanErr, Err: err})
	}

	count, err := res.RowsAffected()
	if err != nil {
		return customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.ExecErr, Err: err})
	}

	if count != 1 {
		return customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.ExecErr, Err: err})
	}

	return nil
}

func (o orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order

	query := `SELECT id, data FROM orders`

	rows, err := o.db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.QueryErr, Err: err})
	}

	for rows.Next() {
		var order models.Order
		err = rows.Scan(&order.ID, &order.Data)
		if err != nil {
			return nil, customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.ScanErr, Err: err})
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, customerr.ErrNormalizer(customerr.ErrorPair{Message: customerr.RowsErr, Err: err})
	}

	return orders, nil
}
