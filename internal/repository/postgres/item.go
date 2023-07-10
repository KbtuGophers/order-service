package postgres

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	db *sqlx.DB
}

func (i *ItemRepository) Select(ctx context.Context) (dest []items.Entity, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemRepository) AddProducts(ctx context.Context, data []items.Entity) (err error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemRepository) DeleteProduct(ctx context.Context, OrderID, ProductID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemRepository) Update(ctx context.Context, data items.Entity) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}
