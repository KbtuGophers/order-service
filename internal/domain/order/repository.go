package order

import (
	"context"
	"github.com/shopspring/decimal"
)

type Repository interface {
	Select(ctx context.Context) (dest []Entity, err error)
	CreateOrder(ctx context.Context, data Entity) (id string, err error)
	Get(ctx context.Context, id string) (dest Entity, err error)
	GetTotal(ctx context.Context, id string) (decimal.Decimal, error)
	Update(ctx context.Context, id string, data Entity) (err error)
	Delete(ctx context.Context, id string) (err error)
}
