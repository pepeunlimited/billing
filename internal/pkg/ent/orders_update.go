// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/billing/internal/pkg/ent/txs"
)

// OrdersUpdate is the builder for updating Orders entities.
type OrdersUpdate struct {
	config
	created_at      *time.Time
	user_id         *int64
	adduser_id      *int64
	txs             map[int]struct{}
	items           map[int]struct{}
	payments        map[int]struct{}
	removedTxs      map[int]struct{}
	removedItems    map[int]struct{}
	clearedPayments bool
	predicates      []predicate.Orders
}

// Where adds a new predicate for the builder.
func (ou *OrdersUpdate) Where(ps ...predicate.Orders) *OrdersUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetCreatedAt sets the created_at field.
func (ou *OrdersUpdate) SetCreatedAt(t time.Time) *OrdersUpdate {
	ou.created_at = &t
	return ou
}

// SetUserID sets the user_id field.
func (ou *OrdersUpdate) SetUserID(i int64) *OrdersUpdate {
	ou.user_id = &i
	ou.adduser_id = nil
	return ou
}

// AddUserID adds i to user_id.
func (ou *OrdersUpdate) AddUserID(i int64) *OrdersUpdate {
	if ou.adduser_id == nil {
		ou.adduser_id = &i
	} else {
		*ou.adduser_id += i
	}
	return ou
}

// AddTxIDs adds the txs edge to Txs by ids.
func (ou *OrdersUpdate) AddTxIDs(ids ...int) *OrdersUpdate {
	if ou.txs == nil {
		ou.txs = make(map[int]struct{})
	}
	for i := range ids {
		ou.txs[ids[i]] = struct{}{}
	}
	return ou
}

// AddTxs adds the txs edges to Txs.
func (ou *OrdersUpdate) AddTxs(t ...*Txs) *OrdersUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ou.AddTxIDs(ids...)
}

// AddItemIDs adds the items edge to Item by ids.
func (ou *OrdersUpdate) AddItemIDs(ids ...int) *OrdersUpdate {
	if ou.items == nil {
		ou.items = make(map[int]struct{})
	}
	for i := range ids {
		ou.items[ids[i]] = struct{}{}
	}
	return ou
}

// AddItems adds the items edges to Item.
func (ou *OrdersUpdate) AddItems(i ...*Item) *OrdersUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ou.AddItemIDs(ids...)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (ou *OrdersUpdate) SetPaymentsID(id int) *OrdersUpdate {
	if ou.payments == nil {
		ou.payments = make(map[int]struct{})
	}
	ou.payments[id] = struct{}{}
	return ou
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (ou *OrdersUpdate) SetNillablePaymentsID(id *int) *OrdersUpdate {
	if id != nil {
		ou = ou.SetPaymentsID(*id)
	}
	return ou
}

// SetPayments sets the payments edge to Payment.
func (ou *OrdersUpdate) SetPayments(p *Payment) *OrdersUpdate {
	return ou.SetPaymentsID(p.ID)
}

// RemoveTxIDs removes the txs edge to Txs by ids.
func (ou *OrdersUpdate) RemoveTxIDs(ids ...int) *OrdersUpdate {
	if ou.removedTxs == nil {
		ou.removedTxs = make(map[int]struct{})
	}
	for i := range ids {
		ou.removedTxs[ids[i]] = struct{}{}
	}
	return ou
}

// RemoveTxs removes txs edges to Txs.
func (ou *OrdersUpdate) RemoveTxs(t ...*Txs) *OrdersUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ou.RemoveTxIDs(ids...)
}

// RemoveItemIDs removes the items edge to Item by ids.
func (ou *OrdersUpdate) RemoveItemIDs(ids ...int) *OrdersUpdate {
	if ou.removedItems == nil {
		ou.removedItems = make(map[int]struct{})
	}
	for i := range ids {
		ou.removedItems[ids[i]] = struct{}{}
	}
	return ou
}

// RemoveItems removes items edges to Item.
func (ou *OrdersUpdate) RemoveItems(i ...*Item) *OrdersUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ou.RemoveItemIDs(ids...)
}

// ClearPayments clears the payments edge to Payment.
func (ou *OrdersUpdate) ClearPayments() *OrdersUpdate {
	ou.clearedPayments = true
	return ou
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OrdersUpdate) Save(ctx context.Context) (int, error) {
	if len(ou.payments) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"payments\"")
	}
	return ou.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrdersUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrdersUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrdersUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrdersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orders.Table,
			Columns: orders.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orders.FieldID,
			},
		},
	}
	if ps := ou.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := ou.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: orders.FieldCreatedAt,
		})
	}
	if value := ou.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: orders.FieldUserID,
		})
	}
	if value := ou.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: orders.FieldUserID,
		})
	}
	if nodes := ou.removedTxs; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.txs; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ou.removedItems; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.items; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.clearedPayments {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.payments; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrdersUpdateOne is the builder for updating a single Orders entity.
type OrdersUpdateOne struct {
	config
	id              int
	created_at      *time.Time
	user_id         *int64
	adduser_id      *int64
	txs             map[int]struct{}
	items           map[int]struct{}
	payments        map[int]struct{}
	removedTxs      map[int]struct{}
	removedItems    map[int]struct{}
	clearedPayments bool
}

// SetCreatedAt sets the created_at field.
func (ouo *OrdersUpdateOne) SetCreatedAt(t time.Time) *OrdersUpdateOne {
	ouo.created_at = &t
	return ouo
}

// SetUserID sets the user_id field.
func (ouo *OrdersUpdateOne) SetUserID(i int64) *OrdersUpdateOne {
	ouo.user_id = &i
	ouo.adduser_id = nil
	return ouo
}

// AddUserID adds i to user_id.
func (ouo *OrdersUpdateOne) AddUserID(i int64) *OrdersUpdateOne {
	if ouo.adduser_id == nil {
		ouo.adduser_id = &i
	} else {
		*ouo.adduser_id += i
	}
	return ouo
}

// AddTxIDs adds the txs edge to Txs by ids.
func (ouo *OrdersUpdateOne) AddTxIDs(ids ...int) *OrdersUpdateOne {
	if ouo.txs == nil {
		ouo.txs = make(map[int]struct{})
	}
	for i := range ids {
		ouo.txs[ids[i]] = struct{}{}
	}
	return ouo
}

// AddTxs adds the txs edges to Txs.
func (ouo *OrdersUpdateOne) AddTxs(t ...*Txs) *OrdersUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ouo.AddTxIDs(ids...)
}

// AddItemIDs adds the items edge to Item by ids.
func (ouo *OrdersUpdateOne) AddItemIDs(ids ...int) *OrdersUpdateOne {
	if ouo.items == nil {
		ouo.items = make(map[int]struct{})
	}
	for i := range ids {
		ouo.items[ids[i]] = struct{}{}
	}
	return ouo
}

// AddItems adds the items edges to Item.
func (ouo *OrdersUpdateOne) AddItems(i ...*Item) *OrdersUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ouo.AddItemIDs(ids...)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (ouo *OrdersUpdateOne) SetPaymentsID(id int) *OrdersUpdateOne {
	if ouo.payments == nil {
		ouo.payments = make(map[int]struct{})
	}
	ouo.payments[id] = struct{}{}
	return ouo
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (ouo *OrdersUpdateOne) SetNillablePaymentsID(id *int) *OrdersUpdateOne {
	if id != nil {
		ouo = ouo.SetPaymentsID(*id)
	}
	return ouo
}

// SetPayments sets the payments edge to Payment.
func (ouo *OrdersUpdateOne) SetPayments(p *Payment) *OrdersUpdateOne {
	return ouo.SetPaymentsID(p.ID)
}

// RemoveTxIDs removes the txs edge to Txs by ids.
func (ouo *OrdersUpdateOne) RemoveTxIDs(ids ...int) *OrdersUpdateOne {
	if ouo.removedTxs == nil {
		ouo.removedTxs = make(map[int]struct{})
	}
	for i := range ids {
		ouo.removedTxs[ids[i]] = struct{}{}
	}
	return ouo
}

// RemoveTxs removes txs edges to Txs.
func (ouo *OrdersUpdateOne) RemoveTxs(t ...*Txs) *OrdersUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ouo.RemoveTxIDs(ids...)
}

// RemoveItemIDs removes the items edge to Item by ids.
func (ouo *OrdersUpdateOne) RemoveItemIDs(ids ...int) *OrdersUpdateOne {
	if ouo.removedItems == nil {
		ouo.removedItems = make(map[int]struct{})
	}
	for i := range ids {
		ouo.removedItems[ids[i]] = struct{}{}
	}
	return ouo
}

// RemoveItems removes items edges to Item.
func (ouo *OrdersUpdateOne) RemoveItems(i ...*Item) *OrdersUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ouo.RemoveItemIDs(ids...)
}

// ClearPayments clears the payments edge to Payment.
func (ouo *OrdersUpdateOne) ClearPayments() *OrdersUpdateOne {
	ouo.clearedPayments = true
	return ouo
}

// Save executes the query and returns the updated entity.
func (ouo *OrdersUpdateOne) Save(ctx context.Context) (*Orders, error) {
	if len(ouo.payments) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"payments\"")
	}
	return ouo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrdersUpdateOne) SaveX(ctx context.Context) *Orders {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OrdersUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrdersUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrdersUpdateOne) sqlSave(ctx context.Context) (o *Orders, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orders.Table,
			Columns: orders.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  ouo.id,
				Type:   field.TypeInt,
				Column: orders.FieldID,
			},
		},
	}
	if value := ouo.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: orders.FieldCreatedAt,
		})
	}
	if value := ouo.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: orders.FieldUserID,
		})
	}
	if value := ouo.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: orders.FieldUserID,
		})
	}
	if nodes := ouo.removedTxs; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.txs; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ouo.removedItems; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.items; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.clearedPayments {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.payments; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	o = &Orders{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}
