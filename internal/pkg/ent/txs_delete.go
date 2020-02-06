// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/billing/internal/pkg/ent/txs"
)

// TxsDelete is the builder for deleting a Txs entity.
type TxsDelete struct {
	config
	predicates []predicate.Txs
}

// Where adds a new predicate to the delete builder.
func (td *TxsDelete) Where(ps ...predicate.Txs) *TxsDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TxsDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TxsDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TxsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: txs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: txs.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TxsDeleteOne is the builder for deleting a single Txs entity.
type TxsDeleteOne struct {
	td *TxsDelete
}

// Exec executes the deletion query.
func (tdo *TxsDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{txs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TxsDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
