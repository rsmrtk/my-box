package amount

type Amount struct {
	Amount          float64 `json:"amount"`
	AmountFormatted string  `json:"amount_formatted"`
	CurrencyCode    string  `json:"currency_code"`
	CurrencySymbol  string  `json:"currency_symbol"`
}
