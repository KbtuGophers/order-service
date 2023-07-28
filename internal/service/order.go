package service

import (
	"context"
	"fmt"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/KbtuGophers/order-service/pkg/payment"
	"github.com/shopspring/decimal"
)

func (s *Service) AddOrder(ctx context.Context, req order.Request) (id string, err error) {
	req.Status = "pending"
	data := order.Entity{
		Status: &req.Status, Data: &req.Data, Currency: &req.Currency, Amount: &req.Amount, CustomerId: &req.CustomerId,
	}

	id, err = s.orderRepository.CreateOrder(ctx, data)

	return
}

func (s *Service) Get(ctx context.Context, id string) (order.Response, error) {
	data, err := s.orderRepository.Get(ctx, id)
	if err != nil {
		return order.Response{}, err
	}

	res := order.ParseFromEntity(data)
	return res, nil
}

func (s *Service) Reorder(ctx context.Context, id string) (string, error) {
	data, err := s.orderRepository.Get(ctx, id)
	if err != nil {
		return "", err
	}

	zero := decimal.Zero
	data.Amount = &zero
	orderID, err := s.orderRepository.CreateOrder(ctx, data)
	if err != nil {
		return "", err
	}

	dest, err := s.itemRepository.Select(ctx, id)
	if err != nil {
		return "", err
	}

	for _, d := range dest {
		req := items.Request{
			OrderID:   orderID,
			StoreID:   *d.StoreID,
			ProductID: *d.ProductID,
			Quantity:  *d.Quantity,
			Price:     *d.Price,
		}
		_, err := s.AddItemToOrder(ctx, req)
		if err != nil {
			return "", err
		}
	}

	return orderID, nil
}

func (s *Service) GetTotalPrice(ctx context.Context, id string) (decimal.Decimal, error) {
	amount, err := s.orderRepository.GetTotal(ctx, id)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return amount, err

}

func (s *Service) GetOrder(ctx context.Context, id string) (order.Response, error) {
	data, err := s.orderRepository.Get(ctx, id)
	if err != nil {
		return order.Response{}, err
	}

	resp := order.ParseFromEntity(data)

	return resp, nil
}

func (s *Service) PayOrder(ctx context.Context, id string) (string, error) {
	orderSrc, err := s.GetOrder(ctx, id)
	if err != nil {
		fmt.Println("get order error")
		return "", err
	}

	req := payment.Request{
		Amount:        orderSrc.Amount,
		Currency:      "KZT",
		AccountId:     orderSrc.CustomerId,
		Language:      "RUS",
		Description:   orderSrc.Data.Info,
		CorrelationId: "",
		Source:        "",
	}

	res, err := s.PaymentClient.CreateBilling(req)
	if err != nil {
		fmt.Println("Create billing error")
		return "", err
	}

	data := order.Entity{BillingID: &res.ID}
	if err := s.orderRepository.Update(ctx, id, data); err != nil {
		return "", err
	}

	redirect, err := s.PaymentClient.Pay(res.Link)
	if err != nil {
		fmt.Println("Pay error")
		return "", err
	}

	err = s.processRepository.PendingToProcessing(ctx, id, "Оформить и оплатить заказ", "Заказ успешно оформлен, ожидается поступление оплаты от клиента")
	if err != nil {
		return "", err
	}

	return redirect, nil
}

func (s *Service) CancelOrder(ctx context.Context, id string) error {
	order, err := s.GetOrder(ctx, id)
	if err != nil {
		fmt.Println("get order error")
		return err
	}

	err = s.PaymentClient.Cancel(order.BillingID)
	if err != nil {
		return err
	}

	err = s.processRepository.ProcessingToCanceled(ctx, id, "Отменить заказ после 20 минут бездействия клиента", "Заказ отменен")

	return nil

}

func (s *Service) ConfirmOrder(ctx context.Context, id string) error {
	err := s.processRepository.ProcessingToCompleted(ctx, id, "Отменить заказ после 20 минут бездействия клиента", "Заказ подтвержден")
	if err != nil {
		return nil
	}

	return nil

}

func (s *Service) GetOrderHistory(ctx context.Context, customerID string) ([]order.Response, error) {
	data, err := s.orderRepository.SelectOrder(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return order.ParseFromEntities(data), nil

}
