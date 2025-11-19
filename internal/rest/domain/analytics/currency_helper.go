package analytics

// GetCurrencySymbol returns the currency symbol for a given currency code
func GetCurrencySymbol(code string) string {
	symbols := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"UAH": "₴",
		"PLN": "zł",
		"CZK": "Kč",
		"CHF": "Fr",
		"JPY": "¥",
		"CNY": "¥",
		"CAD": "C$",
		"AUD": "A$",
		"NZD": "NZ$",
		"SEK": "kr",
		"NOK": "kr",
		"DKK": "kr",
		"RUB": "₽",
		"INR": "₹",
		"BRL": "R$",
		"MXN": "$",
		"ZAR": "R",
		"KRW": "₩",
		"SGD": "S$",
		"HKD": "HK$",
		"TRY": "₺",
		"ILS": "₪",
		"AED": "د.إ",
		"SAR": "﷼",
		"EGP": "E£",
		"THB": "฿",
		"PHP": "₱",
		"IDR": "Rp",
		"MYR": "RM",
		"VND": "₫",
		"BTC": "₿",
		"ETH": "Ξ",
	}

	if symbol, exists := symbols[code]; exists {
		return symbol
	}
	// Return the code itself if no symbol is found
	return code
}

// ValidateCurrency checks if the currency code is valid
func ValidateCurrency(code string) bool {
	validCurrencies := map[string]bool{
		"USD": true, "EUR": true, "GBP": true, "UAH": true, "PLN": true,
		"CZK": true, "CHF": true, "JPY": true, "CNY": true, "CAD": true,
		"AUD": true, "NZD": true, "SEK": true, "NOK": true, "DKK": true,
		"RUB": true, "INR": true, "BRL": true, "MXN": true, "ZAR": true,
		"KRW": true, "SGD": true, "HKD": true, "TRY": true, "ILS": true,
		"AED": true, "SAR": true, "EGP": true, "THB": true, "PHP": true,
		"IDR": true, "MYR": true, "VND": true, "BTC": true, "ETH": true,
	}

	return validCurrencies[code]
}
