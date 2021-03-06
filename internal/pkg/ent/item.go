// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/billing/internal/pkg/ent/item"
	"github.com/pepeunlimited/billing/internal/pkg/ent/orders"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PriceID holds the value of the "price_id" field.
	PriceID int64 `json:"price_id,omitempty"`
	// PlanID holds the value of the "plan_id" field.
	PlanID int64 `json:"plan_id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity uint8 `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges        ItemEdges `json:"edges"`
	orders_items *int
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Orders holds the value of the orders edge.
	Orders *Orders
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) OrdersOrErr() (*Orders, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // price_id
		&sql.NullInt64{}, // plan_id
		&sql.NullInt64{}, // quantity
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Item) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // orders_items
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(values ...interface{}) error {
	if m, n := len(values), len(item.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field price_id", values[0])
	} else if value.Valid {
		i.PriceID = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field plan_id", values[1])
	} else if value.Valid {
		i.PlanID = value.Int64
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quantity", values[2])
	} else if value.Valid {
		i.Quantity = uint8(value.Int64)
	}
	values = values[3:]
	if len(values) == len(item.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field orders_items", value)
		} else if value.Valid {
			i.orders_items = new(int)
			*i.orders_items = int(value.Int64)
		}
	}
	return nil
}

// QueryOrders queries the orders edge of the Item.
func (i *Item) QueryOrders() *OrdersQuery {
	return (&ItemClient{i.config}).QueryOrders(i)
}

// Update returns a builder for updating this Item.
// Note that, you need to call Item.Unwrap() before calling this method, if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", price_id=")
	builder.WriteString(fmt.Sprintf("%v", i.PriceID))
	builder.WriteString(", plan_id=")
	builder.WriteString(fmt.Sprintf("%v", i.PlanID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", i.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
