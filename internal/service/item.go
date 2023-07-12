package service

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/items"
)

func (s *Service) AddItemToOrder(ctx context.Context, request items.Request) (string, error) {
	DefaultCurrency := "KZT"
	data := items.Entity{
		OrderID:   &request.OrderID,
		StoreID:   &request.StoreID,
		ProductID: &request.ProductID,
		Quantity:  &request.Quantity,
		Currency:  &DefaultCurrency,
	}

	id, err := s.itemRepository.AddProducts(ctx, data)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) UpdateQuantity(ctx context.Context, request items.Request) error {
	err := s.itemRepository.UpdateQuantity(ctx, request.OrderID, request.ProductID, request.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) List(ctx context.Context, orderID string) ([]items.Response, error) {
	data, err := s.itemRepository.Select(ctx, orderID)
	if err != nil {
		return nil, err
	}

	res := items.ParseFromEntities(data)
	return res, nil
}
