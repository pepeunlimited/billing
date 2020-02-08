package etc

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"time"
)

func ToOrder(orders *ent.Orders) *orderrpc.Order {
	return &orderrpc.Order{
		Id:        int64(orders.ID),
		CreatedAt: orders.CreatedAt.Format(time.RFC3339),
		UserId:    orders.UserID,
	}
}

func ToOrders(orders []*ent.Orders) []*orderrpc.Order {
	list := make([]*orderrpc.Order, 0)
	for _, order := range orders {
		list = append(list, ToOrder(order))
	}
	return list
}

func ToOrderItems(items []*ent.Item) []*orderrpc.OrderItem {
	return ToOrderItemsWithOrderId(items, 0)
}

func ToOrderItemsWithOrderId(items []*ent.Item, orderId int64) []*orderrpc.OrderItem {
	itemsrpc := make([]*orderrpc.OrderItem, 0)
	for _, item := range items {
		itemsrpc = append(itemsrpc, ToOrderItem(item, orderId))
	}
	return itemsrpc
}

func ToOrderItem(item *ent.Item, orderId int64) *orderrpc.OrderItem {
	if orderId == 0 {
		orderId = int64(item.Edges.Orders.ID)
	}
	return &orderrpc.OrderItem{
		Id:       int64(item.ID),
		PriceId:  item.PriceID,
		Quantity: uint32(item.Quantity),
		OrderId:  orderId,
	}
}

func ToOrderTXsWithOrderId(orderTXs []*ent.Txs, orderId int64) []*orderrpc.OrderTx {
	txs := make([]*orderrpc.OrderTx, 0)
	for _, orderTX := range orderTXs {
		txs = append(txs, ToOrderTX(orderTX, orderId))
	}
	return txs
}

func ToOrderTXs(orderTXs []*ent.Txs) []*orderrpc.OrderTx {
	return ToOrderTXsWithOrderId(orderTXs, 0)
}

func ToOrderTX(orderTX *ent.Txs, orderId int64) *orderrpc.OrderTx {
	if orderId == 0 {
		orderId = int64(orderTX.Edges.Orders.ID)
	}
	return &orderrpc.OrderTx{
		Id:       int64(orderTX.ID),
		OrderId:  orderId,
		Status:   orderrpc.OrderTx_Status(ordersrepo.StatusFromString(orderTX.Status)),
		CreatedAt: orderTX.CreatedAt.Format(time.RFC3339),
	}
}