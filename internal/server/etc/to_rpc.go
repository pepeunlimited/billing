package etc

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/rpc/order"
	"github.com/pepeunlimited/billing/pkg/rpc/payment"
	"time"
)

func ToOrder(orders *ent.Orders) *order.Order {
	return &order.Order{
		Id:        int64(orders.ID),
		CreatedAt: orders.CreatedAt.Format(time.RFC3339),
		UserId:    orders.UserID,
	}
}

func ToOrders(orders []*ent.Orders) []*order.Order {
	list := make([]*order.Order, 0)
	for _, order := range orders {
		list = append(list, ToOrder(order))
	}
	return list
}

func ToOrderItems(items []*ent.Item) []*order.OrderItem {
	return ToOrderItemsWithOrderId(items, 0)
}

func ToOrderItemsWithOrderId(items []*ent.Item, orderId int64) []*order.OrderItem {
	itemsrpc := make([]*order.OrderItem, 0)
	for _, item := range items {
		itemsrpc = append(itemsrpc, ToOrderItem(item, orderId))
	}
	return itemsrpc
}

func ToOrderItem(item *ent.Item, orderId int64) *order.OrderItem {
	if orderId == 0 {
		orderId = int64(item.Edges.Orders.ID)
	}
	return &order.OrderItem{
		Id:       int64(item.ID),
		PriceId:  item.PriceID,
		PlanId:   item.PlanID,
		Quantity: uint32(item.Quantity),
		OrderId:  orderId,
	}
}

func ToOrderTXsWithOrderId(orderTXs []*ent.Txs, orderId int64) []*order.OrderTx {
	txs := make([]*order.OrderTx, 0)
	for _, orderTX := range orderTXs {
		txs = append(txs, ToOrderTX(orderTX, orderId))
	}
	return txs
}

func ToOrderTXs(orderTXs []*ent.Txs) []*order.OrderTx {
	return ToOrderTXsWithOrderId(orderTXs, 0)
}

func ToOrderTX(orderTX *ent.Txs, orderId int64) *order.OrderTx {
	if orderId == 0 {
		orderId = int64(orderTX.Edges.Orders.ID)
	}
	return &order.OrderTx{
		Id:       int64(orderTX.ID),
		OrderId:  orderId,
		Status:   orderTX.Status,
		CreatedAt: orderTX.CreatedAt.Format(time.RFC3339),
	}
}

func ToPaymentInstrument(instrument *ent.Instrument) *payment.PaymentInstrument {
	return &payment.PaymentInstrument{
		Id:   uint32(instrument.ID),
		Type: instrument.Type,
		TypeI18NId: 0,
	}
}

func ToPaymentInstruments(instruments []*ent.Instrument) []*payment.PaymentInstrument {
	list := make([]*payment.PaymentInstrument, 0)
	for _, instrument := range instruments {
		list = append(list, ToPaymentInstrument(instrument))
	}
	return list
}

func ToPayment(from *ent.Payment) *payment.Payment {
	return &payment.Payment{
		Id:					 int64(from.ID),
		PaymentInstrumentId: uint32(from.Edges.Instruments.ID),
		OrderId:           	 int64(from.Edges.Orders.ID),
	}
}

func ToPayments(payments []*ent.Payment) []*payment.Payment {
	list := make([]*payment.Payment, 0)
	for _, payment := range payments {
		list = append(list, ToPayment(payment))
	}
	return list
}