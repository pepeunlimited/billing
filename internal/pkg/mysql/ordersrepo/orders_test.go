package ordersrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"testing"
)

func TestOrdersMySQL_CreateOrder(t *testing.T) {
	ctx := context.TODO()
	ordersrepo := NewOrdersRepository(ent.NewEntClient())
	ordersrepo.Wipe(ctx)
	userID := int64(1)
	items := []*ent.Item{
		&ent.Item{
	 		PriceID:  100,
			Quantity: 1,
		},
		&ent.Item{
			PriceID:  200,
			Quantity: 1,
		},
	}
	order, err := ordersrepo.CreateOrder(ctx, userID, items)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if order.UserID != userID {
		t.FailNow()
	}
	orderItems := order.Edges.Items
 	if len(orderItems) != 2 {
		t.FailNow()
	}
	if orderItems[0].PriceID != 100 {
		t.FailNow()
	}
	if orderItems[1].PriceID != 200 {
		t.FailNow()
	}
}