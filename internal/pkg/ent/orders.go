// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
)

// Orders is the model entity for the Orders schema.
type Orders struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrdersQuery when eager-loading is set.
	Edges          OrdersEdges `json:"edges"`
	payment_orders *int
}

// OrdersEdges holds the relations/edges for other nodes in the graph.
type OrdersEdges struct {
	// Txs holds the value of the txs edge.
	Txs []*Txs
	// Items holds the value of the items edge.
	Items []*Item
	// Payments holds the value of the payments edge.
	Payments *Payment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TxsOrErr returns the Txs value or an error if the edge
// was not loaded in eager-loading.
func (e OrdersEdges) TxsOrErr() ([]*Txs, error) {
	if e.loadedTypes[0] {
		return e.Txs, nil
	}
	return nil, &NotLoadedError{edge: "txs"}
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e OrdersEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[1] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrdersEdges) PaymentsOrErr() (*Payment, error) {
	if e.loadedTypes[2] {
		if e.Payments == nil {
			// The edge payments was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: payment.Label}
		}
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Orders) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // created_at
		&sql.NullInt64{}, // user_id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Orders) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // payment_orders
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Orders fields.
func (o *Orders) assignValues(values ...interface{}) error {
	if m, n := len(values), len(orders.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		o.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[1])
	} else if value.Valid {
		o.UserID = value.Int64
	}
	values = values[2:]
	if len(values) == len(orders.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field payment_orders", value)
		} else if value.Valid {
			o.payment_orders = new(int)
			*o.payment_orders = int(value.Int64)
		}
	}
	return nil
}

// QueryTxs queries the txs edge of the Orders.
func (o *Orders) QueryTxs() *TxsQuery {
	return (&OrdersClient{o.config}).QueryTxs(o)
}

// QueryItems queries the items edge of the Orders.
func (o *Orders) QueryItems() *ItemQuery {
	return (&OrdersClient{o.config}).QueryItems(o)
}

// QueryPayments queries the payments edge of the Orders.
func (o *Orders) QueryPayments() *PaymentQuery {
	return (&OrdersClient{o.config}).QueryPayments(o)
}

// Update returns a builder for updating this Orders.
// Note that, you need to call Orders.Unwrap() before calling this method, if this Orders
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Orders) Update() *OrdersUpdateOne {
	return (&OrdersClient{o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Orders) Unwrap() *Orders {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Orders is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Orders) String() string {
	var builder strings.Builder
	builder.WriteString("Orders(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// OrdersSlice is a parsable slice of Orders.
type OrdersSlice []*Orders

func (o OrdersSlice) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
