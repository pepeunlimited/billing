// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/billing/internal/pkg/ent/plan"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/billing/internal/pkg/ent/subscription"
)

// SubscriptionQuery is the builder for querying Subscription entities.
type SubscriptionQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Subscription
	// eager-loading edges.
	withPlans *PlanQuery
	withFKs   bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (sq *SubscriptionQuery) Where(ps ...predicate.Subscription) *SubscriptionQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SubscriptionQuery) Limit(limit int) *SubscriptionQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SubscriptionQuery) Offset(offset int) *SubscriptionQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *SubscriptionQuery) Order(o ...Order) *SubscriptionQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryPlans chains the current query on the plans edge.
func (sq *SubscriptionQuery) QueryPlans() *PlanQuery {
	query := &PlanQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(subscription.Table, subscription.FieldID, sq.sqlQuery()),
		sqlgraph.To(plan.Table, plan.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, subscription.PlansTable, subscription.PlansColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// First returns the first Subscription entity in the query. Returns *NotFoundError when no subscription was found.
func (sq *SubscriptionQuery) First(ctx context.Context) (*Subscription, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{subscription.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubscriptionQuery) FirstX(ctx context.Context) *Subscription {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Subscription id in the query. Returns *NotFoundError when no id was found.
func (sq *SubscriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subscription.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *SubscriptionQuery) FirstXID(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Subscription entity in the query, returns an error if not exactly one entity was returned.
func (sq *SubscriptionQuery) Only(ctx context.Context) (*Subscription, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{subscription.Label}
	default:
		return nil, &NotSingularError{subscription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubscriptionQuery) OnlyX(ctx context.Context) *Subscription {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Subscription id in the query, returns an error if not exactly one id was returned.
func (sq *SubscriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subscription.Label}
	default:
		err = &NotSingularError{subscription.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (sq *SubscriptionQuery) OnlyXID(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subscriptions.
func (sq *SubscriptionQuery) All(ctx context.Context) ([]*Subscription, error) {
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubscriptionQuery) AllX(ctx context.Context) []*Subscription {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Subscription ids.
func (sq *SubscriptionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(subscription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubscriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubscriptionQuery) Count(ctx context.Context) (int, error) {
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubscriptionQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubscriptionQuery) Exist(ctx context.Context) (bool, error) {
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubscriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubscriptionQuery) Clone() *SubscriptionQuery {
	return &SubscriptionQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]Order{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Subscription{}, sq.predicates...),
		// clone intermediate query.
		sql: sq.sql.Clone(),
	}
}

//  WithPlans tells the query-builder to eager-loads the nodes that are connected to
// the "plans" edge. The optional arguments used to configure the query builder of the edge.
func (sq *SubscriptionQuery) WithPlans(opts ...func(*PlanQuery)) *SubscriptionQuery {
	query := &PlanQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withPlans = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Subscription.Query().
//		GroupBy(subscription.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *SubscriptionQuery) GroupBy(field string, fields ...string) *SubscriptionGroupBy {
	group := &SubscriptionGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = sq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.Subscription.Query().
//		Select(subscription.FieldUserID).
//		Scan(ctx, &v)
//
func (sq *SubscriptionQuery) Select(field string, fields ...string) *SubscriptionSelect {
	selector := &SubscriptionSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = sq.sqlQuery()
	return selector
}

func (sq *SubscriptionQuery) sqlAll(ctx context.Context) ([]*Subscription, error) {
	var (
		nodes       = []*Subscription{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withPlans != nil,
		}
	)
	if sq.withPlans != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Subscription{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withPlans; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Subscription)
		for i := range nodes {
			if fk := nodes[i].plan_subscriptions; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(plan.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "plan_subscriptions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Plans = n
			}
		}
	}

	return nodes, nil
}

func (sq *SubscriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubscriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *SubscriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscription.Table,
			Columns: subscription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscription.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubscriptionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subscription.Table)
	selector := builder.Select(t1.Columns(subscription.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(subscription.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubscriptionGroupBy is the builder for group-by Subscription entities.
type SubscriptionGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubscriptionGroupBy) Aggregate(fns ...Aggregate) *SubscriptionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *SubscriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *SubscriptionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *SubscriptionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SubscriptionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *SubscriptionGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *SubscriptionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SubscriptionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *SubscriptionGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *SubscriptionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SubscriptionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *SubscriptionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *SubscriptionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: SubscriptionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *SubscriptionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *SubscriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SubscriptionGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// SubscriptionSelect is the builder for select fields of Subscription entities.
type SubscriptionSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ss *SubscriptionSelect) Scan(ctx context.Context, v interface{}) error {
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *SubscriptionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *SubscriptionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SubscriptionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *SubscriptionSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *SubscriptionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SubscriptionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *SubscriptionSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *SubscriptionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SubscriptionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *SubscriptionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *SubscriptionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: SubscriptionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *SubscriptionSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *SubscriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *SubscriptionSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}