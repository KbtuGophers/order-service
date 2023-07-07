package order

import (
	"github.com/shopspring/decimal"
)

type Request struct {
	CustomerId string          `json:"customer_id"`
	Amount     decimal.Decimal `json:"amount"`
	Currency   string          `json:"currency"`
	Status     string          `json:"status"`
	Data       Data            `json:"data"`
}

type Response struct {
	ID         string          `json:"id"`
	CustomerId string          `json:"customer_id"`
	Amount     decimal.Decimal `json:"amount"`
	Currency   string          `json:"currency"`
	Status     string          `json:"status"`
	Data       Data            `json:"data"`
}
