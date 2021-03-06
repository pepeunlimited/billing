// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	price_id *int64
	plan_id  *int64
	quantity *uint8
	orders   map[int]struct{}
}

// SetPriceID sets the price_id field.
func (ic *ItemCreate) SetPriceID(i int64) *ItemCreate {
	ic.price_id = &i
	return ic
}

// SetNillablePriceID sets the price_id field if the given value is not nil.
func (ic *ItemCreate) SetNillablePriceID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetPriceID(*i)
	}
	return ic
}

// SetPlanID sets the plan_id field.
func (ic *ItemCreate) SetPlanID(i int64) *ItemCreate {
	ic.plan_id = &i
	return ic
}

// SetNillablePlanID sets the plan_id field if the given value is not nil.
func (ic *ItemCreate) SetNillablePlanID(i *int64) *ItemCreate {
	if i != nil {
		ic.SetPlanID(*i)
	}
	return ic
}

// SetQuantity sets the quantity field.
func (ic *ItemCreate) SetQuantity(u uint8) *ItemCreate {
	ic.quantity = &u
	return ic
}

// SetNillableQuantity sets the quantity field if the given value is not nil.
func (ic *ItemCreate) SetNillableQuantity(u *uint8) *ItemCreate {
	if u != nil {
		ic.SetQuantity(*u)
	}
	return ic
}

// SetOrdersID sets the orders edge to Orders by id.
func (ic *ItemCreate) SetOrdersID(id int) *ItemCreate {
	if ic.orders == nil {
		ic.orders = make(map[int]struct{})
	}
	ic.orders[id] = struct{}{}
	return ic
}

// SetNillableOrdersID sets the orders edge to Orders by id if the given value is not nil.
func (ic *ItemCreate) SetNillableOrdersID(id *int) *ItemCreate {
	if id != nil {
		ic = ic.SetOrdersID(*id)
	}
	return ic
}

// SetOrders sets the orders edge to Orders.
func (ic *ItemCreate) SetOrders(o *Orders) *ItemCreate {
	return ic.SetOrdersID(o.ID)
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	if ic.quantity == nil {
		v := item.DefaultQuantity
		ic.quantity = &v
	}
	if len(ic.orders) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"orders\"")
	}
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	var (
		i     = &Item{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		}
	)
	if value := ic.price_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: item.FieldPriceID,
		})
		i.PriceID = *value
	}
	if value := ic.plan_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: item.FieldPlanID,
		})
		i.PlanID = *value
	}
	if value := ic.quantity; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  *value,
			Column: item.FieldQuantity,
		})
		i.Quantity = *value
	}
	if nodes := ic.orders; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OrdersTable,
			Columns: []string{item.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orders.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	i.ID = int(id)
	return i, nil
}
