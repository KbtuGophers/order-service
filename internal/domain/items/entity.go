package items

import (
	"github.com/shopspring/decimal"
	"time"
)

//created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//id              UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
//order_id        VARCHAR NOT NULL,
//store_id        VARCHAR NOT NULL,
//product_id      VARCHAR NOT NULL,
//quantity        NUMERIC NOT NULL,
//price           NUMERIC NOT NULL,
//currency        VARCHAR NOT NULL,

type Entity struct {
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" db:"updated_at"`
	ID        string           `json:"id" db:"id"`
	OrderID   *string          `json:"order_id" db:"order_id"`
	StoreID   *string          `json:"store_id" db:"store_id"`
	ProductID *string          `json:"product_id" db:"product_id"`
	Quantity  *uint            `json:"quantity" db:"quantity"`
	Price     *decimal.Decimal `json:"price" db:"price"`
	Currency  *string          `json:"currency" db:"currency"`
}
