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

func NewOrderServerValidator() OrderServerValidator {
	return OrderServerValidator{}
}