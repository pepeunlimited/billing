package paymentrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"testing"
)

func TestPaymentMySQL_CreatePaymentInstrument(t *testing.T) {
	ctx := context.TODO()
	paymentrepo := NewPaymentRepository(ent.NewEntClient())
	paymentrepo.Wipe(ctx)
	apple, err := paymentrepo.CreatePaymentInstrument(ctx, PaymentTypeFromString("apple"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if apple.Type != "APPLE" {
		t.FailNow()
	}
	_, err = paymentrepo.CreatePaymentInstrument(ctx, PaymentTypeFromString("apple"))
	if err == nil {
		t.FailNow()
	}
	if err != ErrPaymentInstrumentExist {
		t.FailNow()
	}
	apple, err = paymentrepo.GetPaymentInstrumentByType(ctx, PaymentTypeFromString("apple"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if apple.Type != "APPLE" {
		t.FailNow()
	}
	_, err = paymentrepo.GetPaymentInstrumentByType(ctx, PaymentTypeFromString("google"))
	if err == nil {
		t.FailNow()
	}
	if err != ErrPaymentInstrumentNotExist {
		t.Log(err)
		t.FailNow()
	}
	apple, err = paymentrepo.GetPaymentInstrumentByID(ctx, apple.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	instruments, err := paymentrepo.GetPaymentInstruments(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(instruments) != 1 {
		t.FailNow()
	}
}
