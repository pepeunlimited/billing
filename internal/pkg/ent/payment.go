// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
	"github.com/pepeunlimited/billing/internal/pkg/ent/payment"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges               PaymentEdges `json:"edges"`
	instrument_payments *int
	orders_payments     *int
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// Orders holds the value of the orders edge.
	Orders *Orders
	// Instruments holds the value of the instruments edge.
	Instruments *Instrument
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) OrdersOrErr() (*Orders, error) {
	if e.loadedTypes[0] {
		if e.Orders == nil {
			// The edge orders was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orders.Label}
		}
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// InstrumentsOrErr returns the Instruments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) InstrumentsOrErr() (*Instrument, error) {
	if e.loadedTypes[1] {
		if e.Instruments == nil {
			// The edge instruments was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instrument.Label}
		}
		return e.Instruments, nil
	}
	return nil, &NotLoadedError{edge: "instruments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Payment) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // instrument_payments
		&sql.NullInt64{}, // orders_payments
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(payment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	values = values[0:]
	if len(values) == len(payment.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field instrument_payments", value)
		} else if value.Valid {
			pa.instrument_payments = new(int)
			*pa.instrument_payments = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field orders_payments", value)
		} else if value.Valid {
			pa.orders_payments = new(int)
			*pa.orders_payments = int(value.Int64)
		}
	}
	return nil
}

// QueryOrders queries the orders edge of the Payment.
func (pa *Payment) QueryOrders() *OrdersQuery {
	return (&PaymentClient{pa.config}).QueryOrders(pa)
}

// QueryInstruments queries the instruments edge of the Payment.
func (pa *Payment) QueryInstruments() *InstrumentQuery {
	return (&PaymentClient{pa.config}).QueryInstruments(pa)
}

// Update returns a builder for updating this Payment.
// Note that, you need to call Payment.Unwrap() before calling this method, if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return (&PaymentClient{pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment

func (pa Payments) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
