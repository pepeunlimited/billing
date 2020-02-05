package subscriptionrepo

import "strings"

type Status int

const (
	Unknown = iota + 1
	Paid
	Canceled
	Failed
	Refunded
	Pending
)

func (types Status) String() string {
	return [...]string{"UNKNOWN", "PAID", "CANCELED", "FAILED", "REFUNDED", "PENDING"}[types-1]
}

func SubscriptionTypeFromString(types string) Status {
	types = strings.ToLower(types)
	switch types {
	case "paid":
		return 2
	case "canceled":
		return 4
	case "failed":
		return 4
	case "refunded":
		return 5
	case "pending":
		return 6
	default:
		return 0
	}
}