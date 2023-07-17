package http

import (
	"github.com/KbtuGophers/order-service/internal/config"
	"github.com/KbtuGophers/order-service/internal/domain/items"
	"github.com/KbtuGophers/order-service/internal/domain/order"
	"github.com/KbtuGophers/order-service/internal/service"
	"github.com/KbtuGophers/order-service/pkg/server/status"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
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
	r.Get("/", h.OrderHistory)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Get("/total", h.GetTotal)
		r.Get("/status", h.GetStatus)
		r.Post("/reorder", h.Reorder)
		r.Put("/checkout", h.Checkout)
		r.Put("/cancel", h.Cancel)
		r.Route("/product", func(r chi.Router) {
			r.Post("/", h.AddItem)
			r.Delete("/{product_id}", h.DeleteItem)
			r.Put("/{product_id}", h.ChangeQuantity)
		})

	})
	return r
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	res, err := h.orderService.Get(r.Context(), id)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(res)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return
}

func (h *OrderHandler) GetTotal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	total, err := h.orderService.GetTotalPrice(r.Context(), id)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(map[string]interface{}{
		"total price": total,
	})
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return
}

func (h *OrderHandler) Reorder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	orderID, err := h.orderService.Reorder(r.Context(), id)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(orderID)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return
}

func (h *OrderHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	res, err := h.orderService.GetStatus(r.Context(), id)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(res)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return
}

func (h *OrderHandler) OrderHistory(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("customer_id")
	httpResponse := status.Response{}

	res, err := h.orderService.GetOrderHistory(r.Context(), customerID)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(res)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return

}

func (h *OrderHandler) ChangeQuantity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID := chi.URLParam(r, "product_id")
	quantity := r.URL.Query().Get("quantity")

	httpResponse := status.Response{}

	q, err := strconv.ParseUint(quantity, 10, 64)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	req := items.UpdateRequest{OrderID: id, ProductID: productID, Quantity: uint(q)}

	err = h.orderService.UpdateQuantity(r.Context(), req)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK("changed")
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)
	return

}

func (h *OrderHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpResponse := status.Response{}

	err := h.orderService.CancelOrder(r.Context(), id)
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

	res, err := h.orderService.PayOrder(r.Context(), id)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(res)
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

func (h *OrderHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := items.Request{}
	httpResponse := status.Response{}
	if err := render.Bind(r, &req); err != nil {
		httpResponse = status.BadRequest(err, req)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	req.OrderID = id
	itemID, err := h.orderService.AddItemToOrder(r.Context(), req)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK(itemID)
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)

}

func (h *OrderHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productID := chi.URLParam(r, "product_id")
	httpResponse := status.Response{}

	err := h.orderService.DeleteItem(r.Context(), id, productID)
	if err != nil {
		httpResponse = status.InternalServerError(err)
		httpResponse.Render(w, r)
		render.JSON(w, r, httpResponse)
		return
	}

	httpResponse = status.OK("deleted")
	httpResponse.Render(w, r)
	render.JSON(w, r, httpResponse)

}
