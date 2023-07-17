package processes

import (
	"encoding/json"
	"errors"
	"time"
)

//created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//id              UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
//account_id      VARCHAR NOT NULL,
//order_id        UUID NOT NULL,
//order_status    VARCHAR NOT NULL,
//stage           INTEGER NOT NULL,
//task            VARCHAR NOT NULL,
//method          VARCHAR NOT NULL,
//state           VARCHAR NOT NULL, -- pending/processing/(completed/failed)
//correlation_id  VARCHAR NULL,
//data            JSONB NULL,

type Entity struct {
	CreatedAt     time.Time `json:"-" db:"created_at"`
	UpdatedAt     time.Time `json:"-" db:"updated_at"`
	ID            string    `json:"id" db:"id"`
	AccountID     *string   `json:"account_id" db:"account_id"`
	OrderID       *string   `json:"order_id" db:"order_id"`
	OrderStatus   *string   `json:"order_status" db:"order_status"`
	Stage         *string   `json:"stage" db:"stage"`
	Task          *string   `json:"task" db:"task"`
	Method        *string   `json:"method" db:"method"`
	State         *string   `json:"state" db:"state"`
	CorrelationID *string   `json:"correlation_id" db:"correlation_id"`
	Data          *Data     `json:"data" db:"data"`
}

type Data struct {
	Info string `json:"info"`
}

func (d *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}
