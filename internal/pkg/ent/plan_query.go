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
	"github.com/pepeunlimited/billing/internal/pkg/ent/plan"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/billing/internal/pkg/ent/subscription"
)

// PlanQuery is the builder for querying Plan entities.
type PlanQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Plan
	// eager-loading edges.
	withSubscriptions *SubscriptionQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (pq *PlanQuery) Where(ps ...predicate.Plan) *PlanQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PlanQuery) Limit(limit int) *PlanQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PlanQuery) Offset(offset int) *PlanQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PlanQuery) Order(o ...Order) *PlanQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QuerySubscriptions chains the current query on the subscriptions edge.
func (pq *PlanQuery) QuerySubscriptions() *SubscriptionQuery {
	query := &SubscriptionQuery{config: pq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(plan.Table, plan.FieldID, pq.sqlQuery()),
		sqlgraph.To(subscription.Table, subscription.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, plan.SubscriptionsTable, plan.SubscriptionsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
	return query
}

// First returns the first Plan entity in the query. Returns *NotFoundError when no plan was found.
func (pq *PlanQuery) First(ctx context.Context) (*Plan, error) {
	pls, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pls) == 0 {
		return nil, &NotFoundError{plan.Label}
	}
	return pls[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PlanQuery) FirstX(ctx context.Context) *Plan {
	pl, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pl
}

// FirstID returns the first Plan id in the query. Returns *NotFoundError when no id was found.
func (pq *PlanQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{plan.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PlanQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Plan entity in the query, returns an error if not exactly one entity was returned.
func (pq *PlanQuery) Only(ctx context.Context) (*Plan, error) {
	pls, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pls) {
	case 1:
		return pls[0], nil
	case 0:
		return nil, &NotFoundError{plan.Label}
	default:
		return nil, &NotSingularError{plan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PlanQuery) OnlyX(ctx context.Context) *Plan {
	pl, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pl
}

// OnlyID returns the only Plan id in the query, returns an error if not exactly one id was returned.
func (pq *PlanQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{plan.Label}
	default:
		err = &NotSingularError{plan.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (pq *PlanQuery) OnlyXID(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Plans.
func (pq *PlanQuery) All(ctx context.Context) ([]*Plan, error) {
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PlanQuery) AllX(ctx context.Context) []*Plan {
	pls, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pls
}

// IDs executes the query and returns a list of Plan ids.
func (pq *PlanQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(plan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PlanQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PlanQuery) Count(ctx context.Context) (int, error) {
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PlanQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PlanQuery) Exist(ctx context.Context) (bool, error) {
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PlanQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PlanQuery) Clone() *PlanQuery {
	return &PlanQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]Order{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Plan{}, pq.predicates...),
		// clone intermediate query.
		sql: pq.sql.Clone(),
	}
}

//  WithSubscriptions tells the query-builder to eager-loads the nodes that are connected to
// the "subscriptions" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PlanQuery) WithSubscriptions(opts ...func(*SubscriptionQuery)) *PlanQuery {
	query := &SubscriptionQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withSubscriptions = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TitleI18nID int64 `json:"title_i18n_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Plan.Query().
//		GroupBy(plan.FieldTitleI18nID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PlanQuery) GroupBy(field string, fields ...string) *PlanGroupBy {
	group := &PlanGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = pq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		TitleI18nID int64 `json:"title_i18n_id,omitempty"`
//	}
//
//	client.Plan.Query().
//		Select(plan.FieldTitleI18nID).
//		Scan(ctx, &v)
//
func (pq *PlanQuery) Select(field string, fields ...string) *PlanSelect {
	selector := &PlanSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = pq.sqlQuery()
	return selector
}

func (pq *PlanQuery) sqlAll(ctx context.Context) ([]*Plan, error) {
	var (
		nodes       = []*Plan{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withSubscriptions != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Plan{config: pq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withSubscriptions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Plan)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Subscription(func(s *sql.Selector) {
			s.Where(sql.InValues(plan.SubscriptionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.plan_subscriptions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "plan_subscriptions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "plan_subscriptions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Subscriptions = append(node.Edges.Subscriptions, n)
		}
	}

	return nodes, nil
}

func (pq *PlanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PlanQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PlanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plan.Table,
			Columns: plan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plan.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PlanQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(plan.Table)
	selector := builder.Select(t1.Columns(plan.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(plan.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlanGroupBy is the builder for group-by Plan entities.
type PlanGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PlanGroupBy) Aggregate(fns ...Aggregate) *PlanGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PlanGroupBy) Scan(ctx context.Context, v interface{}) error {
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PlanGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PlanGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PlanGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PlanGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PlanGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PlanGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PlanGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PlanSelect is the builder for select fields of Plan entities.
type PlanSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PlanSelect) Scan(ctx context.Context, v interface{}) error {
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PlanSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PlanSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PlanSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PlanSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PlanSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PlanSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PlanSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PlanSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PlanSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PlanSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PlanSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}