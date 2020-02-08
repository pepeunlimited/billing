package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/pepeunlimited/billing/internal/server/errorz"
	"github.com/pepeunlimited/billing/internal/server/etc"
	"github.com/pepeunlimited/billing/internal/server/validator"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"github.com/twitchtv/twirp"
)

type OrderServer struct {
	orders ordersrepo.OrdersRepository
	valid  validator.OrderServerValidator
	errorz errorz.OrderErrorz
}

func (server OrderServer) CreateOrder(ctx context.Context, params *orderrpc.CreateOrderParams) (*orderrpc.CreateOrderResponse, error) {
	err := server.valid.CreateOrder(params)
	if err != nil {
		return nil, err
	}
	items := etc.FromOrderItems(params.OrderItems)
	order, err := server.orders.CreateOrder(ctx, params.UserId, items)
	if err != nil {
		return nil, twirp.NewError(twirp.Aborted, "create_order_failed")
	}
	return &orderrpc.CreateOrderResponse{
		Order:      etc.ToOrder(order),
		OrderItems: etc.ToOrderItemsWithOrderId(order.Edges.Items, int64(order.ID)),
		OrderTxs:   etc.ToOrderTXsWithOrderId(order.Edges.Txs, int64(order.ID)),
	}, nil
}

func (server OrderServer) GetOrders(ctx context.Context, params *orderrpc.GetOrdersParams) (*orderrpc.GetOrdersResponse, error) {
	err := server.valid.GetOrders(params)
	if err != nil {
		return nil, err
	}
	orders, nextPageToken, err := server.orders.GetOrdersByUserID(ctx, params.UserId, params.PageToken, params.PageSize, false, false, false)

	if err != nil {
		return nil, twirp.InternalError("orders_query_failed")
	}
	return &orderrpc.GetOrdersResponse{
		Orders: etc.ToOrders(orders),
		NextPageToken:nextPageToken,
	}, nil
}

func (server OrderServer) GetOrder(ctx context.Context, params *orderrpc.GetOrderParams) (*orderrpc.Order, error) {
	if err := server.valid.GetOrder(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, false, false)
	if err != nil {
		return nil, server.errorz.IsOrdersError(err)
	}
	return etc.ToOrder(order), nil
}

func (server OrderServer) GetOrderTxs(ctx context.Context, params *orderrpc.GetOrderTxsParams) (*orderrpc.GetOrderTxsResponse, error) {
	if err := server.valid.GetOrderTxs(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, true, false)
	if err != nil {
		return nil, server.errorz.IsOrdersError(err)
	}
	return &orderrpc.GetOrderTxsResponse{OrderTxs: etc.ToOrderTXsWithOrderId(order.Edges.Txs, int64(order.ID))}, nil
}

func (server OrderServer) GetOrderItems(ctx context.Context, params *orderrpc.GetOrderItemsParams) (*orderrpc.GetOrderItemsResponse, error) {
	if err := server.valid.GetOrderItems(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, true, false, false)
	if err != nil {
		return nil, server.errorz.IsOrdersError(err)
	}
	return &orderrpc.GetOrderItemsResponse{OrderItems: etc.ToOrderItemsWithOrderId(order.Edges.Items, int64(order.ID))}, nil
}

func NewOrderServer(client *ent.Client) OrderServer {
	return OrderServer{
		orders:ordersrepo.NewOrdersRepository(client),
		errorz: errorz.NewOrderErrorz(),
	}
}