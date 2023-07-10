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
	Data       struct {
		Info string `json:"info"`
	} `json:"data"`
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
	Data       struct {
		Info string `json:"info"`
	} `json:"data"`
}

func ParseFromEntity(data Entity) Response {
	resp := Response{}
	return resp
}
