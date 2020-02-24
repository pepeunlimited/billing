package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/orders"
	"github.com/pepeunlimited/billing/internal/server/errorz"
	"github.com/pepeunlimited/billing/internal/server/etc"
	"github.com/pepeunlimited/billing/internal/server/validator"
	"github.com/pepeunlimited/billing/pkg/rpc/order"
)

type OrderServer struct {
	orders   orders.OrdersRepository
	valid    validator.OrderServerValidator
}

func (server OrderServer) CreateOrder(ctx context.Context, params *order.CreateOrderParams) (*order.CreateOrderResponse, error) {
	err := server.valid.CreateOrder(params)
	if err != nil {
		return nil, err
	}
	items := etc.FromOrderItems(params.OrderItems)
	fromDB, err := server.orders.CreateOrder(ctx, params.UserId, items)
	if err != nil {
		return nil, errorz.Orders(err)
	}
	return &order.CreateOrderResponse{
		Order:      etc.ToOrder(fromDB),
		OrderItems: etc.ToOrderItemsWithOrderId(fromDB.Edges.Items, int64(fromDB.ID)),
		OrderTxs:   etc.ToOrderTXsWithOrderId(fromDB.Edges.Txs, int64(fromDB.ID)),
	}, nil
}

func (server OrderServer) GetOrders(ctx context.Context, params *order.GetOrdersParams) (*order.GetOrdersResponse, error) {
	err := server.valid.GetOrders(params)
	if err != nil {
		return nil, err
	}
	orders, nextPageToken, err := server.orders.GetOrdersByUserID(ctx, params.UserId, params.PageToken, params.PageSize, false, false, false)
	if err != nil {
		return nil, errorz.Orders(err)
	}
	return &order.GetOrdersResponse{
		Orders: etc.ToOrders(orders),
		NextPageToken:nextPageToken,
	}, nil
}

func (server OrderServer) GetOrder(ctx context.Context, params *order.GetOrderParams) (*order.Order, error) {
	if err := server.valid.GetOrder(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, false, false)
	if err != nil {
		return nil, errorz.Orders(err)
	}
	return etc.ToOrder(order), nil
}

func (server OrderServer) GetOrderTxs(ctx context.Context, params *order.GetOrderTxsParams) (*order.GetOrderTxsResponse, error) {
	if err := server.valid.GetOrderTxs(params); err != nil {
		return nil, err
	}
	fromDB, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, true, false)
	if err != nil {
		return nil, errorz.Orders(err)
	}
	return &order.GetOrderTxsResponse{OrderTxs: etc.ToOrderTXsWithOrderId(fromDB.Edges.Txs, int64(fromDB.ID))}, nil
}

func (server OrderServer) GetOrderItems(ctx context.Context, params *order.GetOrderItemsParams) (*order.GetOrderItemsResponse, error) {
	if err := server.valid.GetOrderItems(params); err != nil {
		return nil, err
	}
	fromDB, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, true, false, false)
	if err != nil {
		return nil, errorz.Orders(err)
	}
	return &order.GetOrderItemsResponse{OrderItems: etc.ToOrderItemsWithOrderId(fromDB.Edges.Items, int64(fromDB.ID))}, nil
}

func NewOrderServer(client *ent.Client) OrderServer {
	return OrderServer{
		orders:   orders.New(client),
	}
}