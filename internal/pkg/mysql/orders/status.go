package orders

import "strings"

type Status int

const (
	Unknown = iota + 1
	Paid
	Canceled
	Failed
	Refunded
	Pending
	Created
)

func (types Status) String() string {
	return [...]string{"UNKNOWN", "PAID", "CANCELED", "FAILED", "REFUNDED", "PENDING", "CREATED"}[types-1]
}

func StatusFromString(types string) Status {
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
	case "created":
		return 7
	default:
		return 1
	}
}