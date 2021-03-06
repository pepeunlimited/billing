// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
)

// InstrumentCreate is the builder for creating a Instrument entity.
type InstrumentCreate struct {
	config
	_type        *string
	type_i18n_id *int64
	payments     map[int]struct{}
}

// SetType sets the type field.
func (ic *InstrumentCreate) SetType(s string) *InstrumentCreate {
	ic._type = &s
	return ic
}

// SetTypeI18nID sets the type_i18n_id field.
func (ic *InstrumentCreate) SetTypeI18nID(i int64) *InstrumentCreate {
	ic.type_i18n_id = &i
	return ic
}

// SetNillableTypeI18nID sets the type_i18n_id field if the given value is not nil.
func (ic *InstrumentCreate) SetNillableTypeI18nID(i *int64) *InstrumentCreate {
	if i != nil {
		ic.SetTypeI18nID(*i)
	}
	return ic
}

// AddPaymentIDs adds the payments edge to Payment by ids.
func (ic *InstrumentCreate) AddPaymentIDs(ids ...int) *InstrumentCreate {
	if ic.payments == nil {
		ic.payments = make(map[int]struct{})
	}
	for i := range ids {
		ic.payments[ids[i]] = struct{}{}
	}
	return ic
}

// AddPayments adds the payments edges to Payment.
func (ic *InstrumentCreate) AddPayments(p ...*Payment) *InstrumentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ic.AddPaymentIDs(ids...)
}

// Save creates the Instrument in the database.
func (ic *InstrumentCreate) Save(ctx context.Context) (*Instrument, error) {
	if ic._type == nil {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if err := instrument.TypeValidator(*ic._type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
	}
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstrumentCreate) SaveX(ctx context.Context) *Instrument {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *InstrumentCreate) sqlSave(ctx context.Context) (*Instrument, error) {
	var (
		i     = &Instrument{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instrument.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instrument.FieldID,
			},
		}
	)
	if value := ic._type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: instrument.FieldType,
		})
		i.Type = *value
	}
	if value := ic.type_i18n_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: instrument.FieldTypeI18nID,
		})
		i.TypeI18nID = *value
	}
	if nodes := ic.payments; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instrument.PaymentsTable,
			Columns: []string{instrument.PaymentsColumn},
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
