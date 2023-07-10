package http

import (
	"github.com/KbtuGophers/order-service/internal/config"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/KbtuGophers/order-service/internal/service"
	"github.com/KbtuGophers/order-service/pkg/server/status"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type OrderHandler struct {
	orderService *service.Service
	config       config.Config
}

func NewOrderHandler(orderService *service.Service, config config.Config) *OrderHandler {
	return &OrderHandler{orderService: orderService, config: config}
}

func (h *OrderHandler) Routes() chi.Router {

	r := chi.NewRouter()
	r.Post("/", h.AddOrder)
	r.Route("/{id}", func(r chi.Router) {
		r.Put("/checkout", h.Checkout)
		r.Put("/cancel", h.Cancel)
	})
	return r
}

func (h *OrderHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	err := h.orderService.PayOrder(r.Context(), id, h.config.ExternalServices.PaymentServiceURL)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(nil)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return

}

func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	err := h.orderService.PayOrder(r.Context(), id, h.config.ExternalServices.PaymentServiceURL)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(nil)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return
}

func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	req := order.Request{}
	httpResponse := status.Response{}
	if err := render.Bind(r, &req); err != nil {
		httpResponse = status.BadRequest(err, req)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	id, err := h.orderService.AddOrder(r.Context(), req)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(id)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)

}
