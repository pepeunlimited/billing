package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"log"
	"testing"
)

func TestOrderServer_GetOrderTxs(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	txs, err := server.GetOrderTxs(ctx, &orderrpc.GetOrderTxsParams{
		UserId:  1,
		OrderId: 1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(txs)
}

func TestOrderServer_GetOrderItems(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	items, err := server.GetOrderItems(ctx, &orderrpc.GetOrderItemsParams{
		UserId:  1,
		OrderId: 1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(items)
}

func TestOrderServer_GetOrder(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	order, err := server.GetOrder(ctx, &orderrpc.GetOrderParams{
		UserId:  1,
		OrderId: 1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(order)
}

func TestOrderServer_GetOrders(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	order, err := server.GetOrders(ctx, &orderrpc.GetOrdersParams{
		UserId:  1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(order)
}

func TestOrderServer_CreateOrder(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	items := []*orderrpc.OrderItem{&orderrpc.OrderItem{
		PriceId:  1,
		Quantity: 1,
	}}
	userId := int64(1)
	order, err := server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: items,
		UserId:     userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if order.OrderItems == nil {
		t.FailNow()
	}
	if order.Order == nil {
		t.FailNow()
	}
	if order.OrderTxs == nil {
		t.FailNow()
	}
	resp, _,_ := server.orders.GetOrdersByUserID(ctx, userId, 0, 20, true, true, true)
	orderIdFromDB := int64(resp[0].ID)
	if orderIdFromDB != order.Order.Id {
		t.FailNow()
	}
	if orderIdFromDB != order.OrderItems[0].OrderId {
		t.FailNow()
	}
	if orderIdFromDB != order.OrderTxs[0].OrderId {
		t.FailNow()
	}
}