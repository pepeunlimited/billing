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
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
	"github.com/pepeunlimited/billing/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/billing/internal/pkg/ent/txs"
)

// OrdersQuery is the builder for querying Orders entities.
type OrdersQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Orders
	// eager-loading edges.
	withTxs      *TxsQuery
	withItems    *ItemQuery
	withPayments *PaymentQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (oq *OrdersQuery) Where(ps ...predicate.Orders) *OrdersQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OrdersQuery) Limit(limit int) *OrdersQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OrdersQuery) Offset(offset int) *OrdersQuery {
	oq.offset = &offset
	return oq
}

// Order adds an order step to the query.
func (oq *OrdersQuery) Order(o ...Order) *OrdersQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryTxs chains the current query on the txs edge.
func (oq *OrdersQuery) QueryTxs() *TxsQuery {
	query := &TxsQuery{config: oq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, oq.sqlQuery()),
		sqlgraph.To(txs.Table, txs.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, orders.TxsTable, orders.TxsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
	return query
}

// QueryItems chains the current query on the items edge.
func (oq *OrdersQuery) QueryItems() *ItemQuery {
	query := &ItemQuery{config: oq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, oq.sqlQuery()),
		sqlgraph.To(item.Table, item.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, orders.ItemsTable, orders.ItemsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
	return query
}

// QueryPayments chains the current query on the payments edge.
func (oq *OrdersQuery) QueryPayments() *PaymentQuery {
	query := &PaymentQuery{config: oq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, oq.sqlQuery()),
		sqlgraph.To(payment.Table, payment.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, orders.PaymentsTable, orders.PaymentsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
	return query
}

// First returns the first Orders entity in the query. Returns *NotFoundError when no orders was found.
func (oq *OrdersQuery) First(ctx context.Context) (*Orders, error) {
	os, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(os) == 0 {
		return nil, &NotFoundError{orders.Label}
	}
	return os[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrdersQuery) FirstX(ctx context.Context) *Orders {
	o, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return o
}

// FirstID returns the first Orders id in the query. Returns *NotFoundError when no id was found.
func (oq *OrdersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orders.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (oq *OrdersQuery) FirstXID(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Orders entity in the query, returns an error if not exactly one entity was returned.
func (oq *OrdersQuery) Only(ctx context.Context) (*Orders, error) {
	os, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(os) {
	case 1:
		return os[0], nil
	case 0:
		return nil, &NotFoundError{orders.Label}
	default:
		return nil, &NotSingularError{orders.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrdersQuery) OnlyX(ctx context.Context) *Orders {
	o, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// OnlyID returns the only Orders id in the query, returns an error if not exactly one id was returned.
func (oq *OrdersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orders.Label}
	default:
		err = &NotSingularError{orders.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (oq *OrdersQuery) OnlyXID(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrdersSlice.
func (oq *OrdersQuery) All(ctx context.Context) ([]*Orders, error) {
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrdersQuery) AllX(ctx context.Context) []*Orders {
	os, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return os
}

// IDs executes the query and returns a list of Orders ids.
func (oq *OrdersQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(orders.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrdersQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrdersQuery) Count(ctx context.Context) (int, error) {
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrdersQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrdersQuery) Exist(ctx context.Context) (bool, error) {
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrdersQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrdersQuery) Clone() *OrdersQuery {
	return &OrdersQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]Order{}, oq.order...),
		unique:     append([]string{}, oq.unique...),
		predicates: append([]predicate.Orders{}, oq.predicates...),
		// clone intermediate query.
		sql: oq.sql.Clone(),
	}
}

//  WithTxs tells the query-builder to eager-loads the nodes that are connected to
// the "txs" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrdersQuery) WithTxs(opts ...func(*TxsQuery)) *OrdersQuery {
	query := &TxsQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withTxs = query
	return oq
}

//  WithItems tells the query-builder to eager-loads the nodes that are connected to
// the "items" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrdersQuery) WithItems(opts ...func(*ItemQuery)) *OrdersQuery {
	query := &ItemQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withItems = query
	return oq
}

//  WithPayments tells the query-builder to eager-loads the nodes that are connected to
// the "payments" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrdersQuery) WithPayments(opts ...func(*PaymentQuery)) *OrdersQuery {
	query := &PaymentQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withPayments = query
	return oq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Orders.Query().
//		GroupBy(orders.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *OrdersQuery) GroupBy(field string, fields ...string) *OrdersGroupBy {
	group := &OrdersGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = oq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Orders.Query().
//		Select(orders.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (oq *OrdersQuery) Select(field string, fields ...string) *OrdersSelect {
	selector := &OrdersSelect{config: oq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = oq.sqlQuery()
	return selector
}

func (oq *OrdersQuery) sqlAll(ctx context.Context) ([]*Orders, error) {
	var (
		nodes       = []*Orders{}
		_spec       = oq.querySpec()
		loadedTypes = [3]bool{
			oq.withTxs != nil,
			oq.withItems != nil,
			oq.withPayments != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Orders{config: oq.config}
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
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oq.withTxs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Orders)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Txs(func(s *sql.Selector) {
			s.Where(sql.InValues(orders.TxsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.orders_txs
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "orders_txs" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "orders_txs" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Txs = append(node.Edges.Txs, n)
		}
	}

	if query := oq.withItems; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Orders)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Item(func(s *sql.Selector) {
			s.Where(sql.InValues(orders.ItemsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.orders_items
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "orders_items" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "orders_items" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Items = append(node.Edges.Items, n)
		}
	}

	if query := oq.withPayments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Orders)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Payment(func(s *sql.Selector) {
			s.Where(sql.InValues(orders.PaymentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.orders_payments
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "orders_payments" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "orders_payments" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Payments = n
		}
	}

	return nodes, nil
}

func (oq *OrdersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrdersQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (oq *OrdersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orders.Table,
			Columns: orders.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orders.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrdersQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(orders.Table)
	selector := builder.Select(t1.Columns(orders.Columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(orders.Columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrdersGroupBy is the builder for group-by Orders entities.
type OrdersGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrdersGroupBy) Aggregate(fns ...Aggregate) *OrdersGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scan the result into the given value.
func (ogb *OrdersGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *OrdersGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrdersGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrdersGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *OrdersGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrdersGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrdersGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *OrdersGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrdersGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrdersGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *OrdersGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrdersGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrdersGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *OrdersGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *OrdersGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogb.sqlQuery().Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OrdersGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql
	columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
	columns = append(columns, ogb.fields...)
	for _, fn := range ogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ogb.fields...)
}

// OrdersSelect is the builder for select fields of Orders entities.
type OrdersSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (os *OrdersSelect) Scan(ctx context.Context, v interface{}) error {
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *OrdersSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (os *OrdersSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrdersSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *OrdersSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (os *OrdersSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrdersSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *OrdersSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (os *OrdersSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrdersSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *OrdersSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (os *OrdersSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrdersSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *OrdersSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *OrdersSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sqlQuery().Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (os *OrdersSelect) sqlQuery() sql.Querier {
	selector := os.sql
	selector.Select(selector.Columns(os.fields...)...)
	return selector
}
