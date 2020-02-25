package validator

import (
	"github.com/pepeunlimited/billing/pkg/rpc/order"
	"github.com/twitchtv/twirp"
)

type OrderServerValidator struct {}

func (v OrderServerValidator) GetOrder(params *order.GetOrderParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrderTxs(params *order.GetOrderTxsParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrderItems(params *order.GetOrderItemsParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v OrderServerValidator) GetOrders(params *order.GetOrdersParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.PageSize == 0 {
		return twirp.RequiredArgumentError("page_size")
	}
	return nil
}

func (v OrderServerValidator) CreateOrder(params *order.CreateOrderParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.OrderItems == nil {
		return twirp.RequiredArgumentError("order_items")
	}
	for _, item := range params.OrderItems {
		if item.Quantity > 255 {
			return twirp.InvalidArgumentError("quantity", "too_large")
		}
		if item.Quantity == 0 {
			return twirp.RequiredArgumentError("quantity")
		}
		if item.PriceId == 0 && item.PlanId == 0 {
			return twirp.RequiredArgumentError("price_id_or_plan_id")
		}
	}
	return nil
}

func NewOrderServerValidator() OrderServerValidator {
	return OrderServerValidator{}
}