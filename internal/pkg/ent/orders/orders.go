// Code generated by entc, DO NOT EDIT.

package orders

const (
	// Label holds the string label denoting the orders type in the database.
	Label = "orders"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUserID holds the string denoting the user_id vertex property in the database.
	FieldUserID = "user_id"

	// Table holds the table name of the orders in the database.
	Table = "orders"
	// TxsTable is the table the holds the txs relation/edge.
	TxsTable = "order_txs"
	// TxsInverseTable is the table name for the Txs entity.
	// It exists in this package in order to avoid circular dependency with the "txs" package.
	TxsInverseTable = "order_txs"
	// TxsColumn is the table column denoting the txs relation/edge.
	TxsColumn = "orders_txs"
	// ItemsTable is the table the holds the items relation/edge.
	ItemsTable = "order_items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "order_items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "orders_items"
)

// Columns holds all SQL columns for orders fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUserID,
}
