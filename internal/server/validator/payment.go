package validator

import (
	paymentrepo "github.com/pepeunlimited/billing/internal/pkg/mysql/payment"
	"github.com/pepeunlimited/billing/pkg/rpc/payment"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

type PaymentServerValidator struct {}

func NewPaymentServerValidator() PaymentServerValidator {
	return PaymentServerValidator{}
}

func (PaymentServerValidator) CreatePayment(params *payment.CreatePaymentParams) error {
	if params.OrderId == 0 {
		return twirp.RequiredArgumentError("order_id")
	}
	if params.PaymentInstrumentId == 0 {
		return twirp.RequiredArgumentError("payment_instrument_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (PaymentServerValidator) CreatePaymentInstrument(params *payment.CreatePaymentInstrumentParams) (paymentrepo.PaymentType, error) {
	types := paymentrepo.PaymentTypeFromString(params.Type)
	if types.String() == "UNKNOWN" {
		return 0, twirp.InvalidArgumentError("type", "unknown_payment_instrument_type")
	}
	return types, nil
}

func (PaymentServerValidator) GetPaymentInstrument(params *payment.GetPaymentInstrumentParams) error {
	if validator.IsEmpty(params.Type) && params.Id == 0 {
		return twirp.RequiredArgumentError("type_or_id")
	}
	if params.Id != 0 {
		return nil
	}
	if paymentrepo.PaymentTypeFromString(params.Type).String() == "UNKNOWN" {
		return twirp.InvalidArgumentError("type", "unknown")
	}
	return nil
}

func (PaymentServerValidator) GetPayments(params *payment.GetPaymentsParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.PageSize == 0 {
		return twirp.RequiredArgumentError("page_size")
	}
	return nil
}

func (PaymentServerValidator) GetPayment(params *payment.GetPaymentParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.OrderId == 0 && params.PaymentId == 0 {
		return twirp.RequiredArgumentError("order_id_or_payment_id")
	}
	return nil
}