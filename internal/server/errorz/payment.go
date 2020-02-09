package errorz

import (
	"github.com/pepeunlimited/billing/internal/pkg/mysql/paymentrepo"
	"github.com/twitchtv/twirp"
	"log"
)

type PaymentErrorz struct {}

func (PaymentErrorz) IsPaymentsError(err error) error {
	switch err {
	case paymentrepo.ErrPaymentInstrumentNotExist:
		return twirp.NotFoundError("payment_instrument_not_found")
	case paymentrepo.ErrPaymentInstrumentExist:
		return twirp.NewError(twirp.AlreadyExists, "payment_instrument")
	case paymentrepo.ErrPaymentNotExist:
		return twirp.NotFoundError("payment_not_found")
	}
	log.Print("payment-server: unknown database issue: "+err.Error())
	return twirp.InternalError(err.Error())
}

func NewPaymentErrorz() PaymentErrorz {
	return PaymentErrorz{}
}