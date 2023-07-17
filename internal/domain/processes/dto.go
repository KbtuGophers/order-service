package processes

type Request struct {
	AccountID     string `json:"account_id"`
	OrderID       string `json:"order_id"`
	OrderStatus   string `json:"order_status"`
	Stage         string `json:"stage"`
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
	Stage         string `json:"stage"`
	Task          string `json:"task"`
	Method        string `json:"method"`
	State         string `json:"state"`
	CorrelationID string `json:"correlation_id"`
	Data          Data   `json:"data"`
}

func ParseFromEntity(data Entity) Response {
	res := Response{
		ID:          data.ID,
		AccountID:   *data.AccountID,
		OrderID:     *data.OrderID,
		OrderStatus: *data.OrderStatus,
		Stage:       *data.Stage,
		Task:        *data.Task,
		Method:      *data.Method,
		State:       *data.State,
	}

	if data.CorrelationID != nil {
		res.Data = *data.Data
	}
	if data.Data != nil {
		res.Data = *data.Data
	}

	return res
}
