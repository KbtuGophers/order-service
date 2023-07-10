package service

import (
	"context"
	"fmt"
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

func (s *Service) GetTotalPrice(ctx context.Context, id string) (decimal.Decimal, error) {
	return decimal.Decimal{}, nil

}

func (s *Service) GetOrder(ctx context.Context, id string) (order.Response, error) {
	data, err := s.orderRepository.Get(ctx, id)
	if err != nil {
		return order.Response{}, err
	}

	resp := order.ParseFromEntity(data)

	return resp, nil
}

func (s *Service) PayOrder(ctx context.Context, id string, url string) error {
	//order, err := s.GetOrder(ctx, id)
	//if err != nil {
	//	fmt.Println("get order error")
	//	return err
	//}

	//total, err := s.GetTotalPrice(ctx, id)
	//if err != nil {
	//	fmt.Println("get total price error")
	//
	//	return err
	//}

	req := payment.Request{
		Amount:        decimal.New(123, 3),
		Currency:      "KZT",
		AccountId:     "1",
		Language:      "RUS",
		Description:   "",
		CorrelationId: "",
		Source:        "",
	}

	res, err := s.PaymentCLinet.CreateBilling(url+"/payment", req)
	if err != nil {
		fmt.Println("Create billing error")
		return err
	}

	err = s.PaymentCLinet.Pay(res.Link)
	if err != nil {
		fmt.Println("Pay error")
		return err
	}

	return nil
}

func (s *Service) CancelOrder(ctx context.Context, id string, url string) error {
	order, err := s.GetOrder(ctx, id)
	if err != nil {
		fmt.Println("get order error")
		return err
	}

	path := url + "/payment/" + order.BillingID + "/cancel"

	err = s.PaymentCLinet.Cancel(path)
	if err != nil {
		return err
	}

	return nil

}
