package processes

import "context"

type Repository interface {
	GetStatus(ctx context.Context, id string) (string, error)
	Checkout(ctx context.Context, id string) error
	Cancel(ctx context.Context, id string) error
	//GetDeliveryStatus(ctx context.Context, id string) error
	//UpdateDelivery(ctx context.Context, id)
}
