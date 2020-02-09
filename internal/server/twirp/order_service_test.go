package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"testing"
)

func TestOrderServer_GetOrderTxs(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	userId := int64(1)
	createOrder,_ := server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: []*orderrpc.OrderItem{&orderrpc.OrderItem{PriceId: 1, Quantity: 1}},
		UserId:     userId,
	})
	txs, err := server.GetOrderTxs(ctx, &orderrpc.GetOrderTxsParams{
		UserId:  userId,
		OrderId: createOrder.Order.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB,_ := server.orders.GetOrderByUserID(ctx, int(createOrder.Order.Id), userId, true, true, true)
	if int64(fromDB.Edges.Txs[0].ID) != txs.OrderTxs[0].Id {
		t.FailNow()
	}
	if fromDB.Edges.Txs[0].Status != txs.OrderTxs[0].Status {
		t.FailNow()
	}
}

func TestOrderServer_GetOrderItems(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	userId := int64(1)
	createOrder,_ := server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: []*orderrpc.OrderItem{&orderrpc.OrderItem{PriceId: 1, Quantity: 1}},
		UserId:     userId,
	})
	items, err := server.GetOrderItems(ctx, &orderrpc.GetOrderItemsParams{
		UserId:  userId,
		OrderId: createOrder.Order.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB,_ := server.orders.GetOrderByUserID(ctx, int(createOrder.Order.Id), userId, true, true, true)
	if int64(fromDB.Edges.Items[0].ID) != items.OrderItems[0].Id {
		t.FailNow()
	}
	if uint32(fromDB.Edges.Items[0].Quantity) != items.OrderItems[0].Quantity {
		t.FailNow()
	}
	if fromDB.Edges.Items[0].PriceID != items.OrderItems[0].PriceId {
		t.FailNow()
	}

}

func TestOrderServer_GetOrder(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	userId := int64(1)
	createOrder,_ := server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: []*orderrpc.OrderItem{&orderrpc.OrderItem{PriceId: 1, Quantity: 1}},
		UserId:     userId,
	})
	order, err := server.GetOrder(ctx, &orderrpc.GetOrderParams{
		UserId:  userId,
		OrderId: createOrder.Order.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if createOrder.Order.Id != order.Id {
		t.FailNow()
	}

}

func TestOrderServer_GetOrders(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	userId := int64(1)
	server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: []*orderrpc.OrderItem{&orderrpc.OrderItem{PriceId: 1, Quantity: 1}},
		UserId:     userId,
	})
	order, err := server.GetOrders(ctx, &orderrpc.GetOrdersParams{
		UserId:  1,
		PageToken:0,
		PageSize:20,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(order.Orders) != 1 {
		t.FailNow()
	}
	if order.NextPageToken == 0 {
		t.FailNow()
	}
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
	fromDB,_ := server.orders.GetOrderByUserID(ctx, int(order.Order.Id), userId, true, true, true)
	orderIdFromDB := int64(fromDB.ID)
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