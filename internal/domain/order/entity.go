package order

import (
	"github.com/shopspring/decimal"
	"time"
)

//created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
//customer_id VARCHAR NOT NULL,
//amount      NUMERIC NOT NULL,
//currency    VARCHAR NOT NULL,
//status      VARCHAR NOT NULL, -- pending/processing/(cancelled/completed)
//data        JSONB NOT NULL

type Entity struct {
	CreatedAt  time.Time        `json:"-" db:"created_at"`
	UpdatedAt  time.Time        `json:"-" db:"updated_at"`
	ID         string           `json:"id" db:"id"`
	CustomerId *string          `json:"customer_id" db:"customer_id"`
	Amount     *decimal.Decimal `json:"amount" db:"amount"`
	Currency   *string          `json:"currency" db:"currency"`
	Status     *string          `json:"status" db:"status"`
	Data       *Data            `json:"data" db:"data"`
}

type Data struct{}
