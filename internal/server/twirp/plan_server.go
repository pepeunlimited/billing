package twirp

import (
	"context"
	"github.com/pepeunlimited/billing/pkg/planrpc"
)

type PlanServer struct {

}

func (p PlanServer) CreatePlan(context.Context, *planrpc.CreatePlanParams) (*planrpc.Plan, error) {
	panic("implement me")
}

func (p PlanServer) UpdatePlan(context.Context, *planrpc.UpdatePlanParams) (*planrpc.Plan, error) {
	panic("implement me")
}

func (p PlanServer) DeletePlan(context.Context, *planrpc.DeletePlanParams) (*planrpc.Plan, error) {
	panic("implement me")
}

func (p PlanServer) GetPlans(context.Context, *planrpc.GetPlansParams) (*planrpc.GetPlansResponse, error) {
	panic("implement me")
}

func (p PlanServer) GetPlan(context.Context, *planrpc.GetPlanParams) (*planrpc.Plan, error) {
	panic("implement me")
}

func NewPlanServer() PlanServer {
	return PlanServer{}
}