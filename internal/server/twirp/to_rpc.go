package twirp

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"time"
)

func toOrder(orders *ent.Orders) *orderrpc.Order {
	return &orderrpc.Order{
		Id:        int64(orders.ID),
		CreatedAt: orders.CreatedAt.Format(time.RFC3339),
		UserId:    orders.UserID,
	}
}

func toOrderItems(items []*ent.Item) []*orderrpc.OrderItem {
	itemsrpc := make([]*orderrpc.OrderItem, 0)
	for _, item := range items {
		itemsrpc = append(itemsrpc, &orderrpc.OrderItem{
			Id:       int64(item.ID),
			PriceId:  item.PriceID,
			Quantity: uint32(item.Quantity),
			OrderId:  int64(item.Edges.Orders.ID),
		})
	}
	return itemsrpc
}

func toOrderTXs(orderTXs []*ent.Txs) []*orderrpc.OrderTx {
	txs := make([]*orderrpc.OrderTx, 0)
	for _, tx := range orderTXs {
		txs = append(txs, &orderrpc.OrderTx{
			Id:       int64(tx.ID),
			OrderId:  int64(tx.Edges.Orders.ID),
			Status:   orderrpc.OrderTx_Status(ordersrepo.StatusFromString(tx.Status)),
			CreatedAt: tx.CreatedAt.Format(time.RFC3339),
		})
	}
	return txs
}
