package warehouse

import "github.com/shopspring/decimal"

type Response struct {
	ID            string          `json:"id"`
	StoreId       string          `json:"store_id"`
	ProductId     string          `json:"product_id"`
	Quantity      uint            `json:"quantity"`
	QuantityMin   uint            `json:"quantity_min"`
	QuantityMax   uint            `json:"quantity_max"`
	Price         decimal.Decimal `json:"price"`
	PriceSpecial  decimal.Decimal `json:"price_special"`
	PricePrevious decimal.Decimal `json:"price_previous"`
	IsAvailable   bool            `json:"is_available"`
}
