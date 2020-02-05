package subscriptionrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"github.com/pepeunlimited/billing/internal/pkg/mysql/planrepo"
	"testing"
	"time"
)

func TestSubscriptionMySQL_Create(t *testing.T) {
	client := ent.NewEntClient()

	ctx := context.TODO()
	subscription := NewSubscriptionRepository(client)
	plans 		 := planrepo.NewPlanRepository(client)
	plans.Delete(ctx)

	// create the plan
	// and assign it to subscriptions

	startAt 	:= time.Now().UTC()
	endAt   	:= time.Now().Add(5 * time.Second).UTC()
	userID  	:= int64(1)
	status  	:= SubscriptionTypeFromString("paid")
	titleI18 	:= int64(1)
	priceID 	:= int64(1)
	length 		:= uint8(1)
	unit  		:= planrepo.PlanUnitFromString("days")

	plan, err := plans.Create(ctx, startAt, endAt, titleI18, priceID, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, plan.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	subscription.Create(ctx, userID, startAt, endAt, status, plan.ID)
}