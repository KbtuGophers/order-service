package items

import "context"

type Repository interface {
	Select(ctx context.Context, OrderID string) (dest []Entity, err error)
	AddProducts(ctx context.Context, data Entity) (id string, err error)
	DeleteProduct(ctx context.Context, OrderID, ProductID string) (err error)
	UpdateQuantity(ctx context.Context, orderID, productID string, quantity uint) (err error)
}
