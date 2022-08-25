package util

const (
	USD = "USD"
	EUR = "EUR"
	RMB = "RMB"
	JPN = "JPN"
)

func IsSupportedCurrency(Currency string) bool {
	switch Currency {
	case USD, EUR, RMB, JPN:
		return true
	}
	return false
}
