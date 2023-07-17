package service

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/processes"
)

func (s *Service) GetStatus(ctx context.Context, orderID string) (processes.Response, error) {
	data, err := s.processRepository.GetStatus(ctx, orderID)
	if err != nil {
		return processes.Response{}, err
	}

	res := processes.ParseFromEntity(data)
	return res, nil
}
