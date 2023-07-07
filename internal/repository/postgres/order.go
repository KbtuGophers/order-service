package postgres

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, req order.Entity) (id string, err error) {
	query := `INSERT INTO orders 
    		  (customer_id, amount, currency, status, data) VALUES 
    		  ($1, $2, $3, $4, $5) RETURNING id`

	args := []any{req.CustomerId, req.Amount, req.Currency, req.Status, req.Data}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

//func (r *OrderRepository) AddItems(req []item.Request) error
