package paymentrepo

import "strings"

type PaymentType int

const (
	Unknown = iota + 1
	Apple
	Google
	Visa
	PayPal
	GiftVoucher
)

func (types PaymentType) String() string {
	return [...]string{"UNKNOWN", "APPLE", "GOOGLE", "VISA", "PAYPAL", "GIFT_VOUCHER"}[types-1]
}

func PaymentTypeFromString(types string) PaymentType {
	types = strings.ToLower(types)
	switch types {
	case "apple":
		return 2
	case "google":
		return 3
	case "visa":
		return 4
	case "paypal":
		return 5
	case "gift_voucher":
		return 6
	default:
		return 1
	}
}