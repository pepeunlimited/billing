package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/pepeunlimited/billing/internal/server/validator"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"github.com/twitchtv/twirp"
	"log"
)

type OrderServer struct {
	orders ordersrepo.OrdersRepository
	valid  validator.OrderServerValidator
}

func (server OrderServer) CreateOrder(ctx context.Context, params *orderrpc.CreateOrderParams) (*orderrpc.CreateOrderResponse, error) {
	err := server.valid.CreateOrder(params)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (server OrderServer) GetOrders(ctx context.Context, params *orderrpc.GetOrdersParams) (*orderrpc.GetOrdersResponse, error) {
	err := server.valid.GetOrders(params)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (server OrderServer) GetOrder(ctx context.Context, params *orderrpc.GetOrderParams) (*orderrpc.Order, error) {
	if err := server.valid.GetOrder(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, false, false)
	if err != nil {
		return nil, server.isOrderError(err)
	}
	return toOrder(order), nil
}

func (server OrderServer) GetOrderTxs(ctx context.Context, params *orderrpc.GetOrderTxsParams) (*orderrpc.GetOrderTxsResponse, error) {
	if err := server.valid.GetOrderTxs(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, true, false)
	if err != nil {
		return nil, server.isOrderError(err)
	}

	return &orderrpc.GetOrderTxsResponse{OrderTxs: toOrderTXs(order.Edges.Txs)}, nil
}

func (server OrderServer) GetOrderItems(ctx context.Context, params *orderrpc.GetOrderItemsParams) (*orderrpc.GetOrderItemsResponse, error) {
	if err := server.valid.GetOrderItems(params); err != nil {
		return nil, err
	}
	order, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, true, false, false)
	if err != nil {
		return nil, server.isOrderError(err)
	}
	return &orderrpc.GetOrderItemsResponse{OrderItems: toOrderItems(order.Edges.Items)}, nil
}

func NewOrderServer(client *ent.Client) OrderServer {
	return OrderServer{
		orders:ordersrepo.NewOrdersRepository(client),
	}
}

func (OrderServer) isOrderError(err error) error {
	switch err {
	case ordersrepo.ErrOrderNotExist:
		return twirp.NotFoundError("order_not_found")
	}
	log.Print("order-server: unknown database issue: "+err.Error())
	return twirp.InternalError(err.Error())
}