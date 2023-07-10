package payment

import "github.com/shopspring/decimal"

type Response struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

type Request struct {
	CorrelationId   string          `json:"correlation_id"`
	Source          string          `json:"source"`
	Amount          decimal.Decimal `json:"amount"`
	Currency        string          `json:"currency"`
	Description     string          `json:"description"`
	TerminalId      string          `json:"terminal_id"`
	AccountId       string          `json:"account_id"`
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Email           string          `json:"email"`
	Language        string          `json:"language"`
	Data            string          `json:"data"`
	CardSave        bool            `json:"card_save"`
	CardId          string          `json:"card_id"`
	BackLink        string          `json:"back_link"`
	FailureBackLink string          `json:"failure_back_link"`
	PostLink        string          `json:"post_link"`
	FailurePostLink string          `json:"failure_post_link"`
	PaymentType     string          `json:"payment_type"`
}

type HandlerResponse struct {
	Success bool     `json:"success"`
	Data    Response `json:"data"`
}
