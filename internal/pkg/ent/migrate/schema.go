// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PlansColumns holds the columns for the "plans" table.
	PlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title_i18n_id", Type: field.TypeInt64},
		{Name: "price_id", Type: field.TypeInt64, Unique: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "length", Type: field.TypeUint8},
		{Name: "unit", Type: field.TypeString, Size: 7},
	}
	// PlansTable holds the schema information for the "plans" table.
	PlansTable = &schema.Table{
		Name:        "plans",
		Columns:     PlansColumns,
		PrimaryKey:  []*schema.Column{PlansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString, Size: 8},
		{Name: "plan_subscriptions", Type: field.TypeInt, Nullable: true},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "subscriptions_plans_subscriptions",
				Columns: []*schema.Column{SubscriptionsColumns[5]},

				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlansTable,
		SubscriptionsTable,
	}
)

func init() {
	SubscriptionsTable.ForeignKeys[0].RefTable = PlansTable
}