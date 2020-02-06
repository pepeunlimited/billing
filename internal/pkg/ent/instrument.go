// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/billing/internal/pkg/ent/instrument"
)

// Instrument is the model entity for the Instrument schema.
type Instrument struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// TypeI18nID holds the value of the "type_i18n_id" field.
	TypeI18nID int64 `json:"type_i18n_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstrumentQuery when eager-loading is set.
	Edges InstrumentEdges `json:"edges"`
}

// InstrumentEdges holds the relations/edges for other nodes in the graph.
type InstrumentEdges struct {
	// Payments holds the value of the payments edge.
	Payments []*Payment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e InstrumentEdges) PaymentsOrErr() ([]*Payment, error) {
	if e.loadedTypes[0] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Instrument) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type
		&sql.NullInt64{},  // type_i18n_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Instrument fields.
func (i *Instrument) assignValues(values ...interface{}) error {
	if m, n := len(values), len(instrument.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[0])
	} else if value.Valid {
		i.Type = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field type_i18n_id", values[1])
	} else if value.Valid {
		i.TypeI18nID = value.Int64
	}
	return nil
}

// QueryPayments queries the payments edge of the Instrument.
func (i *Instrument) QueryPayments() *PaymentQuery {
	return (&InstrumentClient{i.config}).QueryPayments(i)
}

// Update returns a builder for updating this Instrument.
// Note that, you need to call Instrument.Unwrap() before calling this method, if this Instrument
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Instrument) Update() *InstrumentUpdateOne {
	return (&InstrumentClient{i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Instrument) Unwrap() *Instrument {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Instrument is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Instrument) String() string {
	var builder strings.Builder
	builder.WriteString("Instrument(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", type=")
	builder.WriteString(i.Type)
	builder.WriteString(", type_i18n_id=")
	builder.WriteString(fmt.Sprintf("%v", i.TypeI18nID))
	builder.WriteByte(')')
	return builder.String()
}

// Instruments is a parsable slice of Instrument.
type Instruments []*Instrument

func (i Instruments) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
