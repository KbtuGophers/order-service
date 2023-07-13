package service

import (
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/KbtuGophers/order-service/pkg/payment"
	"github.com/KbtuGophers/order-service/pkg/warehouse"
)

type Configuration func(s *Service) error

type Service struct {
	orderRepository order.Repository
	itemRepository  items.Repository
	PaymentClient   *payment.Client
	ProductClient   *warehouse.Client
}

func New(configs ...Configuration) (s *Service, err error) {
	// Create the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithOrderRepository(orderRepository order.Repository, itemRepository items.Repository, paymentClient *payment.Client) Configuration {
	return func(s *Service) error {
		s.orderRepository = orderRepository
		s.PaymentClient = paymentClient
		s.itemRepository = itemRepository
		return nil
	}
}
