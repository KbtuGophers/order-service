package processes

import "context"

type Repository interface {
	GetStatus(ctx context.Context, id string) (Entity, error)
	PendingToProcessing(ctx context.Context, orderID, stage, status string) error
	ProcessingToCompleted(ctx context.Context, orderID, stage, status string) error
	ProcessingToCanceled(ctx context.Context, orderID, stage, status string) error
	//GetDeliveryStatus(ctx context.Context, id string) error
	//UpdateDelivery(ctx context.Context, id)
}
