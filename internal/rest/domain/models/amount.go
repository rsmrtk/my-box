package models

type Amount struct {
	Amount         float64 `json:"amount"`
	CurrencyCode   string  `json:"currency_code"`
	CurrencySymbol string  `json:"currency_symbol"`
}
