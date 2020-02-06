// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
)

// InstrumentQuery is the builder for querying Instrument entities.
type InstrumentQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Instrument
	// eager-loading edges.
	withPayments *PaymentQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (iq *InstrumentQuery) Where(ps ...predicate.Instrument) *InstrumentQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InstrumentQuery) Limit(limit int) *InstrumentQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InstrumentQuery) Offset(offset int) *InstrumentQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *InstrumentQuery) Order(o ...Order) *InstrumentQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryPayments chains the current query on the payments edge.
func (iq *InstrumentQuery) QueryPayments() *PaymentQuery {
	query := &PaymentQuery{config: iq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(instrument.Table, instrument.FieldID, iq.sqlQuery()),
		sqlgraph.To(payment.Table, payment.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, instrument.PaymentsTable, instrument.PaymentsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
	return query
}

// First returns the first Instrument entity in the query. Returns *NotFoundError when no instrument was found.
func (iq *InstrumentQuery) First(ctx context.Context) (*Instrument, error) {
	is, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(is) == 0 {
		return nil, &NotFoundError{instrument.Label}
	}
	return is[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstrumentQuery) FirstX(ctx context.Context) *Instrument {
	i, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return i
}

// FirstID returns the first Instrument id in the query. Returns *NotFoundError when no id was found.
func (iq *InstrumentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instrument.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (iq *InstrumentQuery) FirstXID(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Instrument entity in the query, returns an error if not exactly one entity was returned.
func (iq *InstrumentQuery) Only(ctx context.Context) (*Instrument, error) {
	is, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(is) {
	case 1:
		return is[0], nil
	case 0:
		return nil, &NotFoundError{instrument.Label}
	default:
		return nil, &NotSingularError{instrument.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstrumentQuery) OnlyX(ctx context.Context) *Instrument {
	i, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// OnlyID returns the only Instrument id in the query, returns an error if not exactly one id was returned.
func (iq *InstrumentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instrument.Label}
	default:
		err = &NotSingularError{instrument.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (iq *InstrumentQuery) OnlyXID(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Instruments.
func (iq *InstrumentQuery) All(ctx context.Context) ([]*Instrument, error) {
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstrumentQuery) AllX(ctx context.Context) []*Instrument {
	is, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return is
}

// IDs executes the query and returns a list of Instrument ids.
func (iq *InstrumentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(instrument.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstrumentQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstrumentQuery) Count(ctx context.Context) (int, error) {
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstrumentQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstrumentQuery) Exist(ctx context.Context) (bool, error) {
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstrumentQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstrumentQuery) Clone() *InstrumentQuery {
	return &InstrumentQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]Order{}, iq.order...),
		unique:     append([]string{}, iq.unique...),
		predicates: append([]predicate.Instrument{}, iq.predicates...),
		// clone intermediate query.
		sql: iq.sql.Clone(),
	}
}

//  WithPayments tells the query-builder to eager-loads the nodes that are connected to
// the "payments" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InstrumentQuery) WithPayments(opts ...func(*PaymentQuery)) *InstrumentQuery {
	query := &PaymentQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withPayments = query
	return iq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Instrument.Query().
//		GroupBy(instrument.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InstrumentQuery) GroupBy(field string, fields ...string) *InstrumentGroupBy {
	group := &InstrumentGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = iq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//	}
//
//	client.Instrument.Query().
//		Select(instrument.FieldType).
//		Scan(ctx, &v)
//
func (iq *InstrumentQuery) Select(field string, fields ...string) *InstrumentSelect {
	selector := &InstrumentSelect{config: iq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = iq.sqlQuery()
	return selector
}

func (iq *InstrumentQuery) sqlAll(ctx context.Context) ([]*Instrument, error) {
	var (
		nodes       = []*Instrument{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withPayments != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Instrument{config: iq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withPayments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Instrument)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Payment(func(s *sql.Selector) {
			s.Where(sql.InValues(instrument.PaymentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.instrument_payments
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "instrument_payments" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instrument_payments" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Payments = append(node.Edges.Payments, n)
		}
	}

	return nodes, nil
}

func (iq *InstrumentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstrumentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iq *InstrumentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instrument.Table,
			Columns: instrument.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instrument.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstrumentQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(instrument.Table)
	selector := builder.Select(t1.Columns(instrument.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(instrument.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstrumentGroupBy is the builder for group-by Instrument entities.
type InstrumentGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstrumentGroupBy) Aggregate(fns ...Aggregate) *InstrumentGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scan the result into the given value.
func (igb *InstrumentGroupBy) Scan(ctx context.Context, v interface{}) error {
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *InstrumentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (igb *InstrumentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstrumentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *InstrumentGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (igb *InstrumentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstrumentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *InstrumentGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (igb *InstrumentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstrumentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *InstrumentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (igb *InstrumentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstrumentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *InstrumentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *InstrumentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := igb.sqlQuery().Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InstrumentGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

// InstrumentSelect is the builder for select fields of Instrument entities.
type InstrumentSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (is *InstrumentSelect) Scan(ctx context.Context, v interface{}) error {
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *InstrumentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (is *InstrumentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstrumentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *InstrumentSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (is *InstrumentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstrumentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *InstrumentSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (is *InstrumentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstrumentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *InstrumentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (is *InstrumentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstrumentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *InstrumentSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *InstrumentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *InstrumentSelect) sqlQuery() sql.Querier {
	selector := is.sql
	selector.Select(selector.Columns(is.fields...)...)
	return selector
}