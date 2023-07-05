package service

import (
	"github.com/KbtuGophers/order-service/internal/domain/order"
)

type Configuration func(s *Service) error

type Service struct {
	orderRepository order.Repository
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

func WithOrderRepository(orderRepository order.Repository) Configuration {
	return func(s *Service) error {
		s.orderRepository = orderRepository
		return nil
	}
}
