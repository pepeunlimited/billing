package paymentrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
)

type PaymentRepository interface {

	CreatePayment(ctx context.Context, orderId int, instrumentId int) error


	GetPaymentInstruments()
	CreatePaymentInstruments(ctx context.Context, types PaymentType) (*ent.Instrument, error)
}

type paymentMySQL struct {
	client *ent.Client
}

func (mysql paymentMySQL) CreatePaymentInstruments(ctx context.Context, types PaymentType) (*ent.Instrument, error) {
	panic("implement me")
}

func (mysql paymentMySQL) GetPaymentInstruments() {
	panic("implement me")
}

func (mysql paymentMySQL) CreatePayment(ctx context.Context, orderId int, instrumentId int) error {
	mysql.client.Payment.Create().SetOrdersID(orderId).SetInstrumentsID(instrumentId)
	return nil
}

func NewPaymentRepository(client *ent.Client) PaymentRepository {
	return paymentMySQL{client:client}
}