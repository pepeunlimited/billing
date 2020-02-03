package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/pkg/subscriptionrpc"
)

type SubscriptionServer struct {

}

func (s SubscriptionServer) StartSubscription(context.Context, *subscriptionrpc.CreateSubscriptionParams) (*subscriptionrpc.Subscription, error) {
	panic("implement me")
}

func (s SubscriptionServer) StopSubscription(context.Context, *subscriptionrpc.StopSubscriptionParams) (*subscriptionrpc.Subscription, error) {
	panic("implement me")
}

func (s SubscriptionServer) UpdateSubscription(context.Context, *subscriptionrpc.UpdateSubscriptionParams) (*subscriptionrpc.Subscription, error) {
	panic("implement me")
}

func (s SubscriptionServer) GetSubscription(context.Context, *subscriptionrpc.GetSubscriptionParams) (*subscriptionrpc.Subscription, error) {
	panic("implement me")
}

func (s SubscriptionServer) GetSubscriptions(context.Context, *subscriptionrpc.GetSubscriptionsParams) (*subscriptionrpc.Subscription, error) {
	panic("implement me")
}

func NewSubscriptionServer() SubscriptionServer {
	return SubscriptionServer{}
}