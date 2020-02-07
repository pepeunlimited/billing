// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/pepeunlimited/billing/internal/pkg/ent/migrate"

	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
	"github.com/pepeunlimited/billing/internal/pkg/ent/plan"
	"github.com/pepeunlimited/billing/internal/pkg/ent/subscription"
	"github.com/pepeunlimited/billing/internal/pkg/ent/txs"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Instrument is the client for interacting with the Instrument builders.
	Instrument *InstrumentClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// Orders is the client for interacting with the Orders builders.
	Orders *OrdersClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
	// Plan is the client for interacting with the Plan builders.
	Plan *PlanClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
	// Txs is the client for interacting with the Txs builders.
	Txs *TxsClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:       c,
		Schema:       migrate.NewSchema(c.driver),
		Instrument:   NewInstrumentClient(c),
		Item:         NewItemClient(c),
		Orders:       NewOrdersClient(c),
		Payment:      NewPaymentClient(c),
		Plan:         NewPlanClient(c),
		Subscription: NewSubscriptionClient(c),
		Txs:          NewTxsClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:       cfg,
		Instrument:   NewInstrumentClient(cfg),
		Item:         NewItemClient(cfg),
		Orders:       NewOrdersClient(cfg),
		Payment:      NewPaymentClient(cfg),
		Plan:         NewPlanClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		Txs:          NewTxsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Instrument.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:       cfg,
		Schema:       migrate.NewSchema(cfg.driver),
		Instrument:   NewInstrumentClient(cfg),
		Item:         NewItemClient(cfg),
		Orders:       NewOrdersClient(cfg),
		Payment:      NewPaymentClient(cfg),
		Plan:         NewPlanClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		Txs:          NewTxsClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// InstrumentClient is a client for the Instrument schema.
type InstrumentClient struct {
	config
}

// NewInstrumentClient returns a client for the Instrument from the given config.
func NewInstrumentClient(c config) *InstrumentClient {
	return &InstrumentClient{config: c}
}

// Create returns a create builder for Instrument.
func (c *InstrumentClient) Create() *InstrumentCreate {
	return &InstrumentCreate{config: c.config}
}

// Update returns an update builder for Instrument.
func (c *InstrumentClient) Update() *InstrumentUpdate {
	return &InstrumentUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstrumentClient) UpdateOne(i *Instrument) *InstrumentUpdateOne {
	return c.UpdateOneID(i.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *InstrumentClient) UpdateOneID(id int) *InstrumentUpdateOne {
	return &InstrumentUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Instrument.
func (c *InstrumentClient) Delete() *InstrumentDelete {
	return &InstrumentDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InstrumentClient) DeleteOne(i *Instrument) *InstrumentDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InstrumentClient) DeleteOneID(id int) *InstrumentDeleteOne {
	return &InstrumentDeleteOne{c.Delete().Where(instrument.ID(id))}
}

// Create returns a query builder for Instrument.
func (c *InstrumentClient) Query() *InstrumentQuery {
	return &InstrumentQuery{config: c.config}
}

// Get returns a Instrument entity by its id.
func (c *InstrumentClient) Get(ctx context.Context, id int) (*Instrument, error) {
	return c.Query().Where(instrument.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstrumentClient) GetX(ctx context.Context, id int) *Instrument {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryPayments queries the payments edge of a Instrument.
func (c *InstrumentClient) QueryPayments(i *Instrument) *PaymentQuery {
	query := &PaymentQuery{config: c.config}
	id := i.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(instrument.Table, instrument.FieldID, id),
		sqlgraph.To(payment.Table, payment.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, instrument.PaymentsTable, instrument.PaymentsColumn),
	)
	query.sql = sqlgraph.Neighbors(i.driver.Dialect(), step)

	return query
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	return &ItemCreate{config: c.config}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	return &ItemUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	return c.UpdateOneID(i.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	return &ItemUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	return &ItemDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	return &ItemDeleteOne{c.Delete().Where(item.ID(id))}
}

// Create returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{config: c.config}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryOrders queries the orders edge of a Item.
func (c *ItemClient) QueryOrders(i *Item) *OrdersQuery {
	query := &OrdersQuery{config: c.config}
	id := i.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(item.Table, item.FieldID, id),
		sqlgraph.To(orders.Table, orders.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, item.OrdersTable, item.OrdersColumn),
	)
	query.sql = sqlgraph.Neighbors(i.driver.Dialect(), step)

	return query
}

// OrdersClient is a client for the Orders schema.
type OrdersClient struct {
	config
}

// NewOrdersClient returns a client for the Orders from the given config.
func NewOrdersClient(c config) *OrdersClient {
	return &OrdersClient{config: c}
}

// Create returns a create builder for Orders.
func (c *OrdersClient) Create() *OrdersCreate {
	return &OrdersCreate{config: c.config}
}

// Update returns an update builder for Orders.
func (c *OrdersClient) Update() *OrdersUpdate {
	return &OrdersUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrdersClient) UpdateOne(o *Orders) *OrdersUpdateOne {
	return c.UpdateOneID(o.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *OrdersClient) UpdateOneID(id int) *OrdersUpdateOne {
	return &OrdersUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Orders.
func (c *OrdersClient) Delete() *OrdersDelete {
	return &OrdersDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OrdersClient) DeleteOne(o *Orders) *OrdersDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OrdersClient) DeleteOneID(id int) *OrdersDeleteOne {
	return &OrdersDeleteOne{c.Delete().Where(orders.ID(id))}
}

// Create returns a query builder for Orders.
func (c *OrdersClient) Query() *OrdersQuery {
	return &OrdersQuery{config: c.config}
}

// Get returns a Orders entity by its id.
func (c *OrdersClient) Get(ctx context.Context, id int) (*Orders, error) {
	return c.Query().Where(orders.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrdersClient) GetX(ctx context.Context, id int) *Orders {
	o, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return o
}

// QueryTxs queries the txs edge of a Orders.
func (c *OrdersClient) QueryTxs(o *Orders) *TxsQuery {
	query := &TxsQuery{config: c.config}
	id := o.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, id),
		sqlgraph.To(txs.Table, txs.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, orders.TxsTable, orders.TxsColumn),
	)
	query.sql = sqlgraph.Neighbors(o.driver.Dialect(), step)

	return query
}

// QueryItems queries the items edge of a Orders.
func (c *OrdersClient) QueryItems(o *Orders) *ItemQuery {
	query := &ItemQuery{config: c.config}
	id := o.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, id),
		sqlgraph.To(item.Table, item.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, orders.ItemsTable, orders.ItemsColumn),
	)
	query.sql = sqlgraph.Neighbors(o.driver.Dialect(), step)

	return query
}

// QueryPayments queries the payments edge of a Orders.
func (c *OrdersClient) QueryPayments(o *Orders) *PaymentQuery {
	query := &PaymentQuery{config: c.config}
	id := o.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(orders.Table, orders.FieldID, id),
		sqlgraph.To(payment.Table, payment.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, orders.PaymentsTable, orders.PaymentsColumn),
	)
	query.sql = sqlgraph.Neighbors(o.driver.Dialect(), step)

	return query
}

// PaymentClient is a client for the Payment schema.
type PaymentClient struct {
	config
}

// NewPaymentClient returns a client for the Payment from the given config.
func NewPaymentClient(c config) *PaymentClient {
	return &PaymentClient{config: c}
}

// Create returns a create builder for Payment.
func (c *PaymentClient) Create() *PaymentCreate {
	return &PaymentCreate{config: c.config}
}

// Update returns an update builder for Payment.
func (c *PaymentClient) Update() *PaymentUpdate {
	return &PaymentUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentClient) UpdateOne(pa *Payment) *PaymentUpdateOne {
	return c.UpdateOneID(pa.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentClient) UpdateOneID(id int) *PaymentUpdateOne {
	return &PaymentUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Payment.
func (c *PaymentClient) Delete() *PaymentDelete {
	return &PaymentDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PaymentClient) DeleteOne(pa *Payment) *PaymentDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PaymentClient) DeleteOneID(id int) *PaymentDeleteOne {
	return &PaymentDeleteOne{c.Delete().Where(payment.ID(id))}
}

// Create returns a query builder for Payment.
func (c *PaymentClient) Query() *PaymentQuery {
	return &PaymentQuery{config: c.config}
}

// Get returns a Payment entity by its id.
func (c *PaymentClient) Get(ctx context.Context, id int) (*Payment, error) {
	return c.Query().Where(payment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentClient) GetX(ctx context.Context, id int) *Payment {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QueryOrders queries the orders edge of a Payment.
func (c *PaymentClient) QueryOrders(pa *Payment) *OrdersQuery {
	query := &OrdersQuery{config: c.config}
	id := pa.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(payment.Table, payment.FieldID, id),
		sqlgraph.To(orders.Table, orders.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, payment.OrdersTable, payment.OrdersColumn),
	)
	query.sql = sqlgraph.Neighbors(pa.driver.Dialect(), step)

	return query
}

// QueryInstruments queries the instruments edge of a Payment.
func (c *PaymentClient) QueryInstruments(pa *Payment) *InstrumentQuery {
	query := &InstrumentQuery{config: c.config}
	id := pa.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(payment.Table, payment.FieldID, id),
		sqlgraph.To(instrument.Table, instrument.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, payment.InstrumentsTable, payment.InstrumentsColumn),
	)
	query.sql = sqlgraph.Neighbors(pa.driver.Dialect(), step)

	return query
}

// PlanClient is a client for the Plan schema.
type PlanClient struct {
	config
}

// NewPlanClient returns a client for the Plan from the given config.
func NewPlanClient(c config) *PlanClient {
	return &PlanClient{config: c}
}

// Create returns a create builder for Plan.
func (c *PlanClient) Create() *PlanCreate {
	return &PlanCreate{config: c.config}
}

// Update returns an update builder for Plan.
func (c *PlanClient) Update() *PlanUpdate {
	return &PlanUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanClient) UpdateOne(pl *Plan) *PlanUpdateOne {
	return c.UpdateOneID(pl.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanClient) UpdateOneID(id int) *PlanUpdateOne {
	return &PlanUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Plan.
func (c *PlanClient) Delete() *PlanDelete {
	return &PlanDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlanClient) DeleteOne(pl *Plan) *PlanDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlanClient) DeleteOneID(id int) *PlanDeleteOne {
	return &PlanDeleteOne{c.Delete().Where(plan.ID(id))}
}

// Create returns a query builder for Plan.
func (c *PlanClient) Query() *PlanQuery {
	return &PlanQuery{config: c.config}
}

// Get returns a Plan entity by its id.
func (c *PlanClient) Get(ctx context.Context, id int) (*Plan, error) {
	return c.Query().Where(plan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanClient) GetX(ctx context.Context, id int) *Plan {
	pl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pl
}

// QuerySubscriptions queries the subscriptions edge of a Plan.
func (c *PlanClient) QuerySubscriptions(pl *Plan) *SubscriptionQuery {
	query := &SubscriptionQuery{config: c.config}
	id := pl.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(plan.Table, plan.FieldID, id),
		sqlgraph.To(subscription.Table, subscription.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, plan.SubscriptionsTable, plan.SubscriptionsColumn),
	)
	query.sql = sqlgraph.Neighbors(pl.driver.Dialect(), step)

	return query
}

// SubscriptionClient is a client for the Subscription schema.
type SubscriptionClient struct {
	config
}

// NewSubscriptionClient returns a client for the Subscription from the given config.
func NewSubscriptionClient(c config) *SubscriptionClient {
	return &SubscriptionClient{config: c}
}

// Create returns a create builder for Subscription.
func (c *SubscriptionClient) Create() *SubscriptionCreate {
	return &SubscriptionCreate{config: c.config}
}

// Update returns an update builder for Subscription.
func (c *SubscriptionClient) Update() *SubscriptionUpdate {
	return &SubscriptionUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubscriptionClient) UpdateOne(s *Subscription) *SubscriptionUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SubscriptionClient) UpdateOneID(id int) *SubscriptionUpdateOne {
	return &SubscriptionUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Subscription.
func (c *SubscriptionClient) Delete() *SubscriptionDelete {
	return &SubscriptionDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubscriptionClient) DeleteOne(s *Subscription) *SubscriptionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubscriptionClient) DeleteOneID(id int) *SubscriptionDeleteOne {
	return &SubscriptionDeleteOne{c.Delete().Where(subscription.ID(id))}
}

// Create returns a query builder for Subscription.
func (c *SubscriptionClient) Query() *SubscriptionQuery {
	return &SubscriptionQuery{config: c.config}
}

// Get returns a Subscription entity by its id.
func (c *SubscriptionClient) Get(ctx context.Context, id int) (*Subscription, error) {
	return c.Query().Where(subscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriptionClient) GetX(ctx context.Context, id int) *Subscription {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryPlans queries the plans edge of a Subscription.
func (c *SubscriptionClient) QueryPlans(s *Subscription) *PlanQuery {
	query := &PlanQuery{config: c.config}
	id := s.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(subscription.Table, subscription.FieldID, id),
		sqlgraph.To(plan.Table, plan.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, subscription.PlansTable, subscription.PlansColumn),
	)
	query.sql = sqlgraph.Neighbors(s.driver.Dialect(), step)

	return query
}

// TxsClient is a client for the Txs schema.
type TxsClient struct {
	config
}

// NewTxsClient returns a client for the Txs from the given config.
func NewTxsClient(c config) *TxsClient {
	return &TxsClient{config: c}
}

// Create returns a create builder for Txs.
func (c *TxsClient) Create() *TxsCreate {
	return &TxsCreate{config: c.config}
}

// Update returns an update builder for Txs.
func (c *TxsClient) Update() *TxsUpdate {
	return &TxsUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TxsClient) UpdateOne(t *Txs) *TxsUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TxsClient) UpdateOneID(id int) *TxsUpdateOne {
	return &TxsUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Txs.
func (c *TxsClient) Delete() *TxsDelete {
	return &TxsDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TxsClient) DeleteOne(t *Txs) *TxsDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TxsClient) DeleteOneID(id int) *TxsDeleteOne {
	return &TxsDeleteOne{c.Delete().Where(txs.ID(id))}
}

// Create returns a query builder for Txs.
func (c *TxsClient) Query() *TxsQuery {
	return &TxsQuery{config: c.config}
}

// Get returns a Txs entity by its id.
func (c *TxsClient) Get(ctx context.Context, id int) (*Txs, error) {
	return c.Query().Where(txs.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TxsClient) GetX(ctx context.Context, id int) *Txs {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryOrders queries the orders edge of a Txs.
func (c *TxsClient) QueryOrders(t *Txs) *OrdersQuery {
	query := &OrdersQuery{config: c.config}
	id := t.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(txs.Table, txs.FieldID, id),
		sqlgraph.To(orders.Table, orders.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, txs.OrdersTable, txs.OrdersColumn),
	)
	query.sql = sqlgraph.Neighbors(t.driver.Dialect(), step)

	return query
}
