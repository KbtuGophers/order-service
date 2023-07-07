package processes

type Request struct {
	AccountID     string `json:"account_id"`
	OrderID       string `json:"order_id"`
	OrderStatus   string `json:"order_status"`
	Stage         int    `json:"stage"`
	Task          string `json:"task"`
	Method        string `json:"method"`
	State         string `json:"state"`
	CorrelationID string `json:"correlation_id"`
	Data          Data   `json:"data"`
}

type Response struct {
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	OrderID       string `json:"order_id"`
	OrderStatus   string `json:"order_status"`
	Stage         int    `json:"stage"`
	Task          string `json:"task"`
	Method        string `json:"method"`
	State         string `json:"state"`
	CorrelationID string `json:"correlation_id"`
	Data          Data   `json:"data"`
}
