package errorz

import (
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/twitchtv/twirp"
	"log"
)

type OrderErrorz struct {}

func (OrderErrorz) IsOrdersError(err error) error {
	switch err {
	case ordersrepo.ErrOrderNotExist:
		return twirp.NotFoundError("order_not_found")
	}
	log.Print("order-server: unknown database issue: "+err.Error())
	return twirp.InternalError(err.Error())
}

func NewOrderErrorz() OrderErrorz {
	return OrderErrorz{}
}