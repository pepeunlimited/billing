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