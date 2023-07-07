package items

import "context"

type Repository interface {
	Select(ctx context.Context) (dest []Entity, err error)
	AddProducts(ctx context.Context, data []Entity) (err error)
	DeleteProduct(ctx context.Context, OrderID, ProductID string) (err error)
	Update(ctx context.Context, data Entity) (err error)
}
