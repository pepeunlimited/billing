package etc

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
)

func FromOrderItems(items []*orderrpc.OrderItem) []*ent.Item {
	if items == nil {
		return []*ent.Item{}
	}
	toEnt := make([]*ent.Item, 0)
	for _, item := range items {
		toEnt = append(toEnt, FromOrderItem(item))
	}
	return toEnt
}

func FromOrderItem(item *orderrpc.OrderItem) *ent.Item {
	return &ent.Item{
		PriceID:  item.PriceId,
		Quantity: uint8(item.Quantity),
	}
}