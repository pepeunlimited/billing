package errorz

import (
	"github.com/pepeunlimited/billing/internal/pkg/mysql/orders"
	"github.com/twitchtv/twirp"
	"log"
)

func Orders(err error) error {
	switch err {
	case orders.ErrOrderNotExist:
		return twirp.NotFoundError("order_not_found")
	}
	log.Print("order-server: unknown database issue: "+err.Error())
	return twirp.InternalError(err.Error())
}