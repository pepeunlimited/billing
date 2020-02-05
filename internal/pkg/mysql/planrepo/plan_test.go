package planrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"testing"
	"time"
)

func TestPlanMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	repo := NewPlanRepository(ent.NewEntClient())

	startAt := time.Now()
	endAt 	:= time.Now()
	titleI18nId := int64(2)
	priceId := int64(3)
	length 	:= uint8(2)
	unit 	:= PlanUnitFromString("weeks")

	plan0, err := repo.Create(ctx, startAt, endAt, titleI18nId, priceId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if plan0.Unit == "UNKNOWN" {
		t.FailNow()
	}
}
func TestPlanMySQL_LengthByPlansID(t *testing.T) {
	ctx := context.TODO()
	plans := NewPlanRepository(ent.NewEntClient())
	plans.Delete(ctx)

	startAt := time.Now()
	endAt 	:= time.Now()
	titleI18nId := int64(2)
	priceId := int64(3)
	length 	:= uint8(1)

	// one day
	unit 	:= PlanUnitFromString("days")
	planOneDay, err := plans.Create(ctx, startAt, endAt, titleI18nId, priceId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if planOneDay.Unit == "UNKNOWN" {
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneDay.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,0,1).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one week
	unit 	= PlanUnitFromString("weeks")
	planOneWeek, err := plans.Create(ctx, startAt, endAt, titleI18nId, priceId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneWeek.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,0,7).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one month
	unit 	= PlanUnitFromString("months")
	planOneMonth, err := plans.Create(ctx, startAt, endAt, titleI18nId, priceId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneMonth.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,1,0).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one year
	unit 	= PlanUnitFromString("years")
	planOneYear, err := plans.Create(ctx, startAt, endAt, titleI18nId, priceId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneYear.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(1,0,0).UTC().Equal(endAt) {
		t.FailNow()
	}
}