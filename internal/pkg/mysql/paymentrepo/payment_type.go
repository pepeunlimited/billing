package paymentrepo

import "strings"

type PaymentType int

const (
	Unknown = iota + 1
	Apple
	Google
	Visa
	PayPal
)

func (types PaymentType) String() string {
	return [...]string{"UNKNOWN", "APPLE", "GOOGLE", "Visa", "PayPal"}[types-1]
}

func PaymentTypeFromString(types string) PaymentType {
	types = strings.ToLower(types)
	switch types {
	case "apple":
		return 2
	case "google":
		return 4
	case "visa":
		return 4
	case "paypal":
		return 5
	default:
		return 0
	}
}