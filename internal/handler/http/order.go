package http

import (
	"github.com/KbtuGophers/order-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	orderService *service.Service
}

func NewOrderHandler(orderService *service.Service) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Routes() chi.Router {
	r := chi.NewRouter()

	return r
}
