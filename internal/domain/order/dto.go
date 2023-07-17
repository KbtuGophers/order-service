package order

import (
	"github.com/shopspring/decimal"
	"net/http"
)

type Request struct {
	CustomerId string          `json:"customer_id"`
	Amount     decimal.Decimal `json:"amount"`
	Currency   string          `json:"currency"`
	Status     string          `json:"status"`
	Data       Data            `json:"data"`
}

func (req Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	ID         string          `json:"id"`
	CustomerId string          `json:"customer_id"`
	Amount     decimal.Decimal `json:"amount"`
	Currency   string          `json:"currency"`
	Status     string          `json:"status"`
	BillingID  string          `json:"billing_id"`
	Data       Data            `json:"data"`
}

func ParseFromEntity(data Entity) Response {
	resp := Response{
		ID:         data.ID,
		CustomerId: *data.CustomerId,
		Amount:     *data.Amount,
		Status:     *data.Status,
		Currency:   *data.Currency,
	}

	if data.BillingID != nil {
		resp.BillingID = *data.BillingID
	}
	if data.Data != nil {
		resp.Data = *data.Data
	}

	return resp
}

func ParseFromEntities(data []Entity) []Response {
	resp := make([]Response, len(data))
	for i := range data {
		resp[i] = ParseFromEntity(data[i])
	}

	return resp
}
