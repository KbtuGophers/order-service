package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"strings"
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
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func (r *OrderRepository) GetTotal(ctx context.Context, id string) (amount decimal.Decimal, err error) {
	query := "SELECT amount FROM orders WHERE id=$1"
	args := []any{id}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&amount)

	return
}

func (r *OrderRepository) Update(ctx context.Context, id string, data order.Entity) (err error) {
	sets, args := r.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")
		query := fmt.Sprintf("UPDATE orders SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = r.db.ExecContext(ctx, query, args...)

	}

	return

}

func (r *OrderRepository) prepareArgs(data order.Entity) (sets []string, args []any) {
	if data.BillingID != nil {
		args = append(args, data.BillingID)
		sets = append(sets, fmt.Sprintf("billing_id=$%d", len(args)))
	}
	if data.Amount != nil {
		args = append(args, data.Amount)
		sets = append(sets, fmt.Sprintf("amount=$%d", len(args)))
	}
	if data.Currency != nil {
		args = append(args, data.Currency)
		sets = append(sets, fmt.Sprintf("currency=$%d", len(args)))
	}

	return
}

func (r *OrderRepository) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM orders WHERE id=$1"
	args := []any{id}

	_, err = r.db.ExecContext(ctx, query, args...)
	return
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, data order.Entity) (id string, err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}

	query := `INSERT INTO orders 
    		  (customer_id, amount, currency, status, data) VALUES 
    		  ($1, $2, $3, $4, $5) RETURNING id`

	dataByte, err := json.Marshal(data.Data)
	if err != nil {
		return
	}
	args := []any{data.CustomerId, data.Amount, data.Currency, data.Status, string(dataByte)}

	err = tx.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		tx.Rollback()
		return
	}

	query = `INSERT INTO processes (account_id, order_id, order_status, stage, task, state, method)
						VALUES ($1, $2, $3, $4, $5, $6,$7)`
	args = []any{data.CustomerId, id, "Заказ успешно создан, ожидается оформление от клиента", "Создание заказа", "order", "pending", "create"}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r *OrderRepository) SelectOrder(ctx context.Context, customerID string) (res []order.Entity, err error) {
	query := "SELECT * FROM orders WHERE customer_id=$1 ORDER BY updated_at desc"
	args := []any{customerID}

	err = r.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return
	}

	return
}
