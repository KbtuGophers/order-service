package postgres

import (
	"context"
	"fmt"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
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

func (i *ItemRepository) AddProducts(ctx context.Context, data items.Entity) (string, error) {
	tx, err := i.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}

	price := *data.Price
	quantity := *data.Quantity
	query := `UPDATE orders SET amount = amount + $1 WHERE id = $2`
	args := []any{price.Mul(decimal.New(int64(quantity), 0)), *data.OrderID}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	query = `
 		INSERT INTO items (order_id, store_id, product_id, quantity, price, currency)
    			    VALUES ($1, $2, $3, $4, $5, $6) 
 		RETURNING id
    `
	args = []any{data.OrderID, data.StoreID, data.ProductID, data.Quantity, data.Price, data.Currency}

	var id string
	err = tx.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return id, nil
}

func (i *ItemRepository) DeleteProduct(ctx context.Context, orderID, productID string) (err error) {
	tx, err := i.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	query := `SELECT * FROM items WHERE order_id=$1 and product_id=$2`
	args := []any{orderID, productID}

	item := items.Entity{}
	err = tx.GetContext(ctx, &item, query, args...)
	if err != nil {
		tx.Rollback()
		return
	}

	price := *item.Price
	quantity := *item.Quantity
	query = `UPDATE orders SET amount = amount - $1 WHERE id = $2`
	args = []any{price.Mul(decimal.New(int64(quantity), 0)), *item.OrderID}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	// delete item
	query = `
   		DELETE FROM items WHERE order_id=$1 and product_id=$2
    `

	args = []any{orderID, productID}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return
}

func (i *ItemRepository) GetItemPrice(ctx context.Context, orderID, productID string) (string, error) {
	fmt.Println(orderID, productID)
	query := `SELECT price FROM items WHERE order_id=$1 and product_id=$2`
	args := []any{orderID, productID}
	var price string
	err := i.db.QueryRowContext(ctx, query, args...).Scan(&price)
	if err != nil {
		return "", err
	}

	return price, nil

}

func (i *ItemRepository) UpdateQuantity(ctx context.Context, data items.Entity) (err error) {
	tx, err := i.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}
	fmt.Println(data.OrderID, data.ProductID)
	query := `UPDATE items SET quantity=$1 WHERE order_id=$2 and product_id=$3`
	args := []any{data.Quantity, data.OrderID, data.ProductID}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return
	}

	price := *data.Price
	quantity := *data.Quantity
	query = `UPDATE orders SET amount = $1 WHERE id = $2`
	args = []any{price.Mul(decimal.New(int64(quantity), 0)), data.OrderID}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}
