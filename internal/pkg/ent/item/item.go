// Code generated by entc, DO NOT EDIT.

package item

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent/schema"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPriceID holds the string denoting the price_id vertex property in the database.
	FieldPriceID = "price_id"
	// FieldPlanID holds the string denoting the plan_id vertex property in the database.
	FieldPlanID = "plan_id"
	// FieldQuantity holds the string denoting the quantity vertex property in the database.
	FieldQuantity = "quantity"

	// Table holds the table name of the item in the database.
	Table = "order_items"
	// OrdersTable is the table the holds the orders relation/edge.
	OrdersTable = "order_items"
	// OrdersInverseTable is the table name for the Orders entity.
	// It exists in this package in order to avoid circular dependency with the "orders" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "orders_items"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldPriceID,
	FieldPlanID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Item type.
var ForeignKeys = []string{
	"orders_items",
}

var (
	fields = schema.Item{}.Fields()

	// descQuantity is the schema descriptor for quantity field.
	descQuantity = fields[2].Descriptor()
	// DefaultQuantity holds the default value on creation for the quantity field.
	DefaultQuantity = descQuantity.Default.(uint8)
)
