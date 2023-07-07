package service

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/order"
)

func (s *Service) AddOrder(ctx context.Context, req order.Request) (id string, err error) {
	data := order.Entity{
		Status: &req.Status, Data: &req.Data, Currency: &req.Currency, Amount: &req.Amount, CustomerId: &req.CustomerId,
	}

	id, err = s.orderRepository.CreateOrder(ctx, data)

	return
}
