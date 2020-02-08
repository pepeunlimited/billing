package paymentrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
	"log"
	"time"
)

var (
	ErrPaymentInstrumentExist = errors.New("payments: instrument exist")
	ErrPaymentInstrumentNotExist = errors.New("payments: instrument not exist")
)

type PaymentRepository interface {
	CreatePaymentInstrument(ctx context.Context, types PaymentType) (*ent.Instrument, error)
	CreatePayment(ctx context.Context, orderId int, instrumentId int) (*ent.Payment, error)

	GetPaymentInstruments(ctx context.Context) ([]*ent.Instrument, error)
	GetPaymentInstrumentByType(ctx context.Context, types PaymentType) (*ent.Instrument, error)
	GetPaymentInstrumentByID(ctx context.Context, id int) (*ent.Instrument, error)

	Wipe(ctx context.Context)
}

type paymentMySQL struct {
	client *ent.Client
}

func (mysql paymentMySQL) GetPaymentInstrumentByID(ctx context.Context, id int) (*ent.Instrument, error) {
	i, err := mysql.client.Instrument.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPaymentInstrumentNotExist
		}
		return nil, err
	}
	return i, nil
}

func (mysql paymentMySQL) GetPaymentInstruments(ctx context.Context) ([]*ent.Instrument, error) {
	return mysql.client.Instrument.Query().All(ctx)
}

func (mysql paymentMySQL) GetPaymentInstrumentByType(ctx context.Context, types PaymentType) (*ent.Instrument, error) {
	i, err := mysql.client.Instrument.Query().Where(instrument.Type(types.String())).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPaymentInstrumentNotExist
		}
		return nil, err
	}
	return i, nil
}

func (mysql paymentMySQL) Wipe(ctx context.Context) {
	mysql.client.Txs.Delete().ExecX(ctx)
	mysql.client.Item.Delete().ExecX(ctx)
	mysql.client.Payment.Delete().ExecX(ctx)
	mysql.client.Orders.Delete().ExecX(ctx)
	mysql.client.Instrument.Delete().ExecX(ctx)

}

func (mysql paymentMySQL) CreatePaymentInstrument(ctx context.Context, types PaymentType) (*ent.Instrument, error) {
	instrument, err := mysql.client.Instrument.Create().SetType(types.String()).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrPaymentInstrumentExist
		}
		return nil, err
	}
	return instrument, nil
}

func (mysql paymentMySQL) CreatePayment(ctx context.Context, orderId int, instrumentId int) (*ent.Payment, error) {
	if _, err := mysql.GetPaymentInstrumentByID(ctx, instrumentId); err != nil {
		return nil, err
	}
	tx, err := mysql.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	pay, err := tx.Payment.Create().SetOrdersID(orderId).SetInstrumentsID(instrumentId).Save(ctx)
	if err != nil {
		mysql.rollback(tx)
		return nil, err
	}
	_, err = tx.Txs.Create().SetOrdersID(orderId).SetStatus("PAID").SetCreatedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		mysql.rollback(tx)
		return nil, err
	}
	if err := mysql.commit(tx); err != nil {
		return nil, err
	}
	return pay, nil
}

func NewPaymentRepository(client *ent.Client) PaymentRepository {
	return paymentMySQL{client:client}
}

func (mysql paymentMySQL) rollback(tx *ent.Tx){
	if err := tx.Rollback(); err != nil {
		log.Print("payments: failed execute rollback.."+err.Error())
	}
}

func (mysql paymentMySQL) commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		log.Print("payments: failed execute commit.."+err.Error())
		return err
	}
	return nil
}