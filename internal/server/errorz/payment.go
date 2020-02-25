package errorz

import (
	"github.com/pepeunlimited/billing/internal/pkg/mysql/payment"
	"github.com/twitchtv/twirp"
	"log"
)

func Payment(err error) error {
	switch err {
	case payment.ErrPaymentInstrumentNotExist:
		return twirp.NotFoundError("payment_instrument_not_found")
	case payment.ErrPaymentInstrumentExist:
		return twirp.NewError(twirp.AlreadyExists, "payment_instrument")
	case payment.ErrPaymentNotExist:
		return twirp.NotFoundError("payment_not_found")
	case payment.ErrPaymentExist:
		return twirp.NewError(twirp.AlreadyExists, "payment")
	}
	log.Print("payment-server: unknown database issue: "+err.Error())
	return twirp.InternalError(err.Error())
}