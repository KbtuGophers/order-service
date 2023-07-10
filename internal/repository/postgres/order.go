package postgres

import (
	"context"
	"encoding/json"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type OrderRepository struct {
	db *sqlx.DB
}

func (r *OrderRepository) Select(ctx context.Context) (dest []order.Entity, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *OrderRepository) Get(ctx context.Context, id string) (dest order.Entity, err error) {
	query := "SELECT * FROM orders WHERE id=$1"
	args := []any{id}

	err = r.db.GetContext(ctx, &dest, query, args...)
	return
}

func (r *OrderRepository) GetTotal(ctx context.Context, id string) (decimal.Decimal, error) {
	//query := `
	//	SELECT SUM()
	//`
	return decimal.Decimal{}, nil
}

func (r *OrderRepository) Update(ctx context.Context, id string, data order.Entity) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *OrderRepository) Delete(ctx context.Context, id string) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, data order.Entity) (id string, err error) {
	query := `INSERT INTO orders 
    		  (customer_id, amount, currency, status, data) VALUES 
    		  ($1, $2, $3, $4, $5) RETURNING id`

	dataByte, err := json.Marshal(data.Data)
	if err != nil {
		return
	}
	args := []any{data.CustomerId, data.Amount, data.Currency, data.Status, string(dataByte)}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

//func (r *OrderRepository) AddItems(req []item.Request) error
