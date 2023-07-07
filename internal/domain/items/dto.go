package items

import "github.com/shopspring/decimal"

type Request struct {
	OrderID   string `json:"order_id"`
	StoreID   string `json:"store_id"`
	ProductID string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}

type Response struct {
	ID        string          `json:"id"`
	OrderID   string          `json:"order_id"`
	StoreID   string          `json:"store_id"`
	ProductID string          `json:"product_id"`
	Quantity  uint            `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	Currency  string          `json:"currency"`
}
