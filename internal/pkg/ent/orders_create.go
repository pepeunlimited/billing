// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
	"github.com/pepeunlimited/billing/internal/pkg/ent/txs"
)

// OrdersCreate is the builder for creating a Orders entity.
type OrdersCreate struct {
	config
	created_at *time.Time
	user_id    *int64
	txs        map[int]struct{}
	items      map[int]struct{}
	payments   map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (oc *OrdersCreate) SetCreatedAt(t time.Time) *OrdersCreate {
	oc.created_at = &t
	return oc
}

// SetUserID sets the user_id field.
func (oc *OrdersCreate) SetUserID(i int64) *OrdersCreate {
	oc.user_id = &i
	return oc
}

// AddTxIDs adds the txs edge to Txs by ids.
func (oc *OrdersCreate) AddTxIDs(ids ...int) *OrdersCreate {
	if oc.txs == nil {
		oc.txs = make(map[int]struct{})
	}
	for i := range ids {
		oc.txs[ids[i]] = struct{}{}
	}
	return oc
}

// AddTxs adds the txs edges to Txs.
func (oc *OrdersCreate) AddTxs(t ...*Txs) *OrdersCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oc.AddTxIDs(ids...)
}

// AddItemIDs adds the items edge to Item by ids.
func (oc *OrdersCreate) AddItemIDs(ids ...int) *OrdersCreate {
	if oc.items == nil {
		oc.items = make(map[int]struct{})
	}
	for i := range ids {
		oc.items[ids[i]] = struct{}{}
	}
	return oc
}

// AddItems adds the items edges to Item.
func (oc *OrdersCreate) AddItems(i ...*Item) *OrdersCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return oc.AddItemIDs(ids...)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (oc *OrdersCreate) SetPaymentsID(id int) *OrdersCreate {
	if oc.payments == nil {
		oc.payments = make(map[int]struct{})
	}
	oc.payments[id] = struct{}{}
	return oc
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (oc *OrdersCreate) SetNillablePaymentsID(id *int) *OrdersCreate {
	if id != nil {
		oc = oc.SetPaymentsID(*id)
	}
	return oc
}

// SetPayments sets the payments edge to Payment.
func (oc *OrdersCreate) SetPayments(p *Payment) *OrdersCreate {
	return oc.SetPaymentsID(p.ID)
}

// Save creates the Orders in the database.
func (oc *OrdersCreate) Save(ctx context.Context) (*Orders, error) {
	if oc.created_at == nil {
		return nil, errors.New("ent: missing required field \"created_at\"")
	}
	if oc.user_id == nil {
		return nil, errors.New("ent: missing required field \"user_id\"")
	}
	if len(oc.payments) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"payments\"")
	}
	return oc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrdersCreate) SaveX(ctx context.Context) *Orders {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OrdersCreate) sqlSave(ctx context.Context) (*Orders, error) {
	var (
		o     = &Orders{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orders.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orders.FieldID,
			},
		}
	)
	if value := oc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: orders.FieldCreatedAt,
		})
		o.CreatedAt = *value
	}
	if value := oc.user_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: orders.FieldUserID,
		})
		o.UserID = *value
	}
	if nodes := oc.txs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orders.TxsTable,
			Columns: []string{orders.TxsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: txs.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.items; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orders.ItemsTable,
			Columns: []string{orders.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.payments; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   orders.PaymentsTable,
			Columns: []string{orders.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}
