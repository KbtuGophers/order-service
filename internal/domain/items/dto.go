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

func ParseFromEntity(data Entity) Response {
	res := Response{}
	res.ID = data.ID
	res.OrderID = *data.OrderID
	if data.ProductID != nil {
		res.ProductID = *data.ProductID
	}
	if data.Currency != nil {
		res.Currency = *data.Currency
	}
	if data.StoreID != nil {
		res.StoreID = *data.StoreID
	}

	return res

}

func ParseFromEntities(data []Entity) []Response {
	res := make([]Response, len(data))

	for i := range data {
		res[i] = ParseFromEntity(data[i])
	}

	return res

}
