package validator

import (
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"github.com/twitchtv/twirp"
)

type OrderServerValidator struct {}

func (v OrderServerValidator) GetOrder(params *orderrpc.GetOrderParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrderTxs(params *orderrpc.GetOrderTxsParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrderItems(params *orderrpc.GetOrderItemsParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrders(params *orderrpc.GetOrdersParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.PageSize == 0 {
		return twirp.RequiredArgumentError("page_size")
	}
	if params.PageToken == 0 {
		return twirp.RequiredArgumentError("page_token")
	}
	return nil
}

func (v OrderServerValidator) CreateOrder(params *orderrpc.CreateOrderParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	for _, item := range params.OrderItems {
		if item.Quantity == 0 {
			return twirp.RequiredArgumentError("quantity")
		}
		if item.PriceId == 0 {
			return twirp.RequiredArgumentError("price_id")
		}
	}
	return nil
}

func NewOrderServerValidator() OrderServerValidator {
	return OrderServerValidator{}
}