package etc

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/rpc/order"
)

func FromOrderItems(items []*order.OrderItem) []*ent.Item {
	if items == nil {
		return []*ent.Item{}
	}
	toEnt := make([]*ent.Item, 0)
	for _, item := range items {
		toEnt = append(toEnt, FromOrderItem(item))
	}
	return toEnt
}

func FromOrderItem(item *order.OrderItem) *ent.Item {
	return &ent.Item{
		PriceID:  item.PriceId,
		PlanID:   item.PlanId,
		Quantity: uint8(item.Quantity),
	}
}