package service

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/shopspring/decimal"
	"strconv"
)

func (s *Service) AddItemToOrder(ctx context.Context, request items.Request) (string, error) {
	DefaultCurrency := "KZT"
	data := items.Entity{
		OrderID:   &request.OrderID,
		StoreID:   &request.StoreID,
		ProductID: &request.ProductID,
		Quantity:  &request.Quantity,
		Currency:  &DefaultCurrency,
		Price:     &request.Price,
	}

	id, err := s.itemRepository.AddProducts(ctx, data)
	if err != nil {
		return "", err
	}

	return id, nil
}

func ConvertToDecimal(val string) (decimal.Decimal, error) {
	valINT, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return decimal.New(valINT, 0), nil
}

func (s *Service) UpdateQuantity(ctx context.Context, request items.UpdateRequest) error {
	price, err := s.itemRepository.GetItemPrice(ctx, request.OrderID, request.ProductID)
	if err != nil {
		return err
	}
	priceDec, err := ConvertToDecimal(price)
	if err != nil {
		return err
	}

	data := items.Entity{OrderID: &request.OrderID, ProductID: &request.ProductID, Quantity: &request.Quantity, Price: &priceDec}
	err = s.itemRepository.UpdateQuantity(ctx, data)
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

func (s *Service) DeleteItem(ctx context.Context, orderID, productID string) error {
	err := s.itemRepository.DeleteProduct(ctx, orderID, productID)
	if err != nil {
		return err
	}

	return nil
}
