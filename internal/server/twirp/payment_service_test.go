package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/pkg/orderrpc"
	"github.com/pepeunlimited/billing/pkg/paymentrpc"
	"log"
	"testing"
)

func TestPaymentServer_CreatePayment(t *testing.T) {
	ctx := context.TODO()
	server := NewOrderServer(ent.NewEntClient())
	paymentServer := NewPaymentServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	userId := int64(1)
	createOrder,_ := server.CreateOrder(ctx, &orderrpc.CreateOrderParams{
		OrderItems: []*orderrpc.OrderItem{&orderrpc.OrderItem{PriceId: 1, Quantity: 1}},
		UserId:     userId,
	})
	instrument, err := paymentServer.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "apple",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	payment, err := paymentServer.CreatePayment(ctx, &paymentrpc.CreatePaymentParams{
		OrderId:             createOrder.Order.Id,
		PaymentInstrumentId: instrument.Id,
		UserId:              userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err := paymentServer.payment.GetPaymentByOrderID(ctx, int(createOrder.Order.Id))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if uint32(fromDB.Edges.Instruments.ID) != payment.PaymentInstrumentId {
		t.FailNow()
	}
	if int64(fromDB.Edges.Orders.ID) != payment.OrderId {
		t.FailNow()
	}

	getPayment, err := paymentServer.GetPayment(ctx, &paymentrpc.GetPaymentParams{
		PaymentId: payment.Id,
		UserId:    userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if getPayment.Id != int64(fromDB.ID) {
		t.FailNow()
	}
	getPayment2, err := paymentServer.GetPayment(ctx, &paymentrpc.GetPaymentParams{
		OrderId: createOrder.Order.Id,
		UserId:    userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if getPayment2.Id != int64(fromDB.ID) {
		t.FailNow()
	}

	payments, err := paymentServer.GetPayments(ctx, &paymentrpc.GetPaymentsParams{
		UserId:    userId,
		PageSize:  20,
		PageToken: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(payments.Payments) != 1 {
		t.FailNow()
	}
	if payments.NextPageToken == 0 {
		t.FailNow()
	}
	if payments.Payments[0].Id != int64(fromDB.ID) {
		t.FailNow()
	}
}

func TestPaymentServer_CreatePaymentInstrumentAndGetByID(t *testing.T) {
	ctx := context.TODO()
	server := NewPaymentServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	instrument, err := server.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "apple",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	paymentInstrument, err := server.GetPaymentInstrument(ctx, &paymentrpc.GetPaymentInstrumentParams{
		Id: instrument.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if paymentInstrument.Id != instrument.Id {
		t.FailNow()
	}
}

func TestPaymentServer_CreatePaymentInstrumentAndGetByType(t *testing.T) {
	ctx := context.TODO()
	server := NewPaymentServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	instrument, err := server.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "apple",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	paymentInstrument, err := server.GetPaymentInstrument(ctx, &paymentrpc.GetPaymentInstrumentParams{
		Type:"apple",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if paymentInstrument.Id != instrument.Id {
		t.FailNow()
	}
}

func TestPaymentServer_GetPaymentInstruments(t *testing.T) {
	ctx := context.TODO()
	server := NewPaymentServer(ent.NewEntClient())
	server.orders.Wipe(ctx)
	server.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "apple",
	})
	server.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "google",
	})
	server.CreatePaymentInstrument(ctx, &paymentrpc.CreatePaymentInstrumentParams{
		Type: "gift_voucher",
	})
	instruments, err := server.payment.GetPaymentInstruments(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(instruments)
}