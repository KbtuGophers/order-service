package postgres

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	db *sqlx.DB
}

func (i *ItemRepository) Select(ctx context.Context, OrderID string) (dest []items.Entity, err error) {
	query := `SELECT * FROM items WHERE order_id=$1`
	args := []any{OrderID}

	err = i.db.SelectContext(ctx, &dest, query, args...)

	return
}

func (i *ItemRepository) AddProducts(ctx context.Context, data items.Entity) (id string, err error) {
	query := `
 		INSERT INTO items (order_id, store_id, product_id, quantity, price, currency)
    			    VALUES ($1, $2, $3, $4, $5, $6) 
 		RETURNING id
    `
	args := []any{data.OrderID, data.StoreID, data.ProductID, data.Quantity, data.Price, data.Currency}
	err = i.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (i *ItemRepository) DeleteProduct(ctx context.Context, orderID, productID string) (err error) {
	query := `
   		DELETE FROM items WHERE order_id=$1 and product_id=$2
    `
	args := []any{orderID, productID}

	_, err = i.db.ExecContext(ctx, query, args...)

	return
}

func (i *ItemRepository) UpdateQuantity(ctx context.Context, orderID, productID string, quantity uint) (err error) {

	query := `UPDATE items SET quantity=$1 WHERE order_id=$2 and product_id=$3`
	args := []any{quantity, orderID, productID}

	_, err = i.db.ExecContext(ctx, query, args...)

	return
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}
