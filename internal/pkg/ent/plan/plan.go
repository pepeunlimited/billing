// Code generated by entc, DO NOT EDIT.

package plan

import (
	"github.com/pepeunlimited/billing/internal/pkg/ent/schema"
)

const (
	// Label holds the string label denoting the plan type in the database.
	Label = "plan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitleI18nID holds the string denoting the title_i18n_id vertex property in the database.
	FieldTitleI18nID = "title_i18n_id"
	// FieldPriceID holds the string denoting the price_id vertex property in the database.
	FieldPriceID = "price_id"
	// FieldStartAt holds the string denoting the start_at vertex property in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at vertex property in the database.
	FieldEndAt = "end_at"
	// FieldLength holds the string denoting the length vertex property in the database.
	FieldLength = "length"
	// FieldUnit holds the string denoting the unit vertex property in the database.
	FieldUnit = "unit"

	// Table holds the table name of the plan in the database.
	Table = "plans"
	// SubscriptionsTable is the table the holds the subscriptions relation/edge.
	SubscriptionsTable = "subscriptions"
	// SubscriptionsInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionsInverseTable = "subscriptions"
	// SubscriptionsColumn is the table column denoting the subscriptions relation/edge.
	SubscriptionsColumn = "plan_subscriptions"
)

// Columns holds all SQL columns for plan fields.
var Columns = []string{
	FieldID,
	FieldTitleI18nID,
	FieldPriceID,
	FieldStartAt,
	FieldEndAt,
	FieldLength,
	FieldUnit,
}

var (
	fields = schema.Plan{}.Fields()

	// descUnit is the schema descriptor for unit field.
	descUnit = fields[5].Descriptor()
	// UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	UnitValidator = descUnit.Validators[0].(func(string) error)
)