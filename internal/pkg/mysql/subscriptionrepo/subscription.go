package subscriptionrepo

import (
	"context"
	"github.com/pepeunlimited/billing/internal/pkg/ent"
	"time"
)


type SubscriptionRepository interface {
	Create(ctx context.Context, userID int64, startAt time.Time, endAt time.Time, status Status, planID int)
	GetByID(ctx context.Context)
	GetByUserID(ctx context.Context)
	Delete(ctx context.Context)
	Update(ctx context.Context)
}

type subscriptionMySQL struct {
	client *ent.Client
}

func (mysql subscriptionMySQL) Create(ctx context.Context, userID int64, startAt time.Time, endAt time.Time, status Status, planID int) {
	mysql.
		client.
		Subscription.
		Create().
		SetUserID(userID).
		SetEndAt(endAt.UTC()).
		SetStartAt(startAt.UTC()).
		SetStatus(status.String()).
		SetPlansID(planID).
		SaveX(ctx)
}

func (mysql subscriptionMySQL) GetByID(ctx context.Context) {
	panic("implement me")
}

func (mysql subscriptionMySQL) GetByUserID(ctx context.Context) {
	panic("implement me")
}

func (mysql subscriptionMySQL) Delete(ctx context.Context) {
	panic("implement me")
}

func (mysql subscriptionMySQL) Update(ctx context.Context) {
	panic("implement me")
}

func NewSubscriptionRepository(client *ent.Client) SubscriptionRepository {
	return subscriptionMySQL{client:client}
}