package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/ordersrepo"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/paymentrepo"
	"github.com/pepeunlimited/billing/internal/server/errorz"
	"github.com/pepeunlimited/billing/internal/server/etc"
	"github.com/pepeunlimited/billing/internal/server/validator"
	"github.com/pepeunlimited/billing/pkg/paymentrpc"
)

type PaymentServer struct {
	payment paymentrepo.PaymentRepository
	orders 	ordersrepo.OrdersRepository
	valid  	validator.PaymentServerValidator
	paymenterr errorz.PaymentErrorz
	orderserr errorz.OrderErrorz
}

func (server PaymentServer) CreatePayment(ctx context.Context, params *paymentrpc.CreatePaymentParams) (*paymentrpc.Payment, error) {
	if err := server.valid.CreatePayment(params); err != nil {
		return nil, err
	}
	_, err := server.orders.GetOrderByUserID(ctx, int(params.OrderId), params.UserId, false, false, false)
	if err != nil {
		return nil, server.orderserr.IsOrdersError(err)
	}
	payment, err := server.payment.CreatePayment(ctx, int(params.OrderId), int(params.PaymentInstrumentId))
	if err != nil {
		return nil, server.paymenterr.IsPaymentsError(err)
	}
	return etc.ToPayment(payment), nil
}

func (server PaymentServer) GetPayments(ctx context.Context, params *paymentrpc.GetPaymentsParams) (*paymentrpc.GetPaymentsResponse, error) {
	err := server.valid.GetPayments(params)
	if err != nil {
		return nil, err
	}
	payments, nextPageToken, err := server.payment.GetPayments(ctx, params.UserId, params.PageToken, params.PageSize)
	return &paymentrpc.GetPaymentsResponse{
		Payments:      etc.ToPayments(payments),
		NextPageToken: nextPageToken,
	}, nil
}

func (server PaymentServer) GetPayment(ctx context.Context, params *paymentrpc.GetPaymentParams) (*paymentrpc.Payment, error) {
	err := server.valid.GetPayment(params)
	if err != nil {
		return nil, err
	}
	var payment *ent.Payment
	if params.OrderId == 0 {
		payment, err = server.payment.GetPaymentByID(ctx, int(params.PaymentId))
	} else {
		payment, err = server.payment.GetPaymentByOrderID(ctx, int(params.OrderId))
	}
	if err != nil {
		return nil, server.paymenterr.IsPaymentsError(err)
	}
	return etc.ToPayment(payment), nil
}

func (server PaymentServer) CreatePaymentInstrument(ctx context.Context, params *paymentrpc.CreatePaymentInstrumentParams) (*paymentrpc.PaymentInstrument, error) {
	instrument, err := server.valid.CreatePaymentInstrument(params)
	if err != nil {
		return nil, err
	}
	paymentInstrument, err := server.payment.CreatePaymentInstrument(ctx, instrument)
	if err != nil {
		return nil, server.paymenterr.IsPaymentsError(err)
	}
	return etc.ToPaymentInstrument(paymentInstrument), nil
}

func (server PaymentServer) GetPaymentInstrument(ctx context.Context, params *paymentrpc.GetPaymentInstrumentParams) (*paymentrpc.PaymentInstrument, error) {
	err := server.valid.GetPaymentInstrument(params)
	if err != nil {
		return nil, err
	}
	var instrument *ent.Instrument
	if params.Id == 0 {
		// query by type
		instrument, err = server.payment.GetPaymentInstrumentByType(ctx, paymentrepo.PaymentTypeFromString(params.Type))
	} else {
		// query by id
		instrument, err = server.payment.GetPaymentInstrumentByID(ctx, int(params.Id))
	}
	if err != nil {
		return nil, server.paymenterr.IsPaymentsError(err)
	}
	return etc.ToPaymentInstrument(instrument), nil
}

func (server PaymentServer) GetPaymentInstruments(ctx context.Context, params *paymentrpc.GetPaymentInstrumentsParams) (*paymentrpc.GetPaymentInstrumentsResponse, error) {
	instruments, err := server.payment.GetPaymentInstruments(ctx)
	if err != nil {
		return nil, server.paymenterr.IsPaymentsError(err)
	}
	return &paymentrpc.GetPaymentInstrumentsResponse{PaymentInstruments: etc.ToPaymentInstruments(instruments)}, nil
}

func NewPaymentServer(client *ent.Client) PaymentServer {
	return PaymentServer{
		payment:		paymentrepo.NewPaymentRepository(client),
		valid:			validator.NewPaymentServerValidator(),
		orders:			ordersrepo.NewOrdersRepository(client),
		orderserr:		errorz.NewOrderErrorz(),
		paymenterr: 	errorz.NewPaymentErrorz(),
	}
}