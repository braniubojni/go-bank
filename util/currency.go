package util

import "slices"

const (
	USD = "USD"
	EUR = "EUR"
	PLN = "PLN"
)

var currencies = []string{
	USD,
	EUR,
	PLN,
}

func IsSupportedCurrency(currency string) bool {
	if slices.Contains(currencies, currency) {
		return true
	}
	return false
}
