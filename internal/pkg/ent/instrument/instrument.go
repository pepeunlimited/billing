// Code generated by entc, DO NOT EDIT.

package instrument

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent/schema"
)

const (
	// Label holds the string label denoting the instrument type in the database.
	Label = "instrument"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type vertex property in the database.
	FieldType = "type"
	// FieldTypeI18nID holds the string denoting the type_i18n_id vertex property in the database.
	FieldTypeI18nID = "type_i18n_id"

	// Table holds the table name of the instrument in the database.
	Table = "payment_instruments"
	// PaymentsTable is the table the holds the payments relation/edge.
	PaymentsTable = "payments"
	// PaymentsInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentsInverseTable = "payments"
	// PaymentsColumn is the table column denoting the payments relation/edge.
	PaymentsColumn = "instrument_payments"
)

// Columns holds all SQL columns for instrument fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldTypeI18nID,
}

var (
	fields = schema.Instrument{}.Fields()

	// descType is the schema descriptor for type field.
	descType = fields[0].Descriptor()
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator = descType.Validators[0].(func(string) error)
)
