package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// => `order_items`
type Item struct {
	ent.Schema
}

func (Item) Config() ent.Config {
	return ent.Config{Table:"order_items"}
}

func (Item) Fields() []ent.Field {
	return []ent.Field {
		field.Int64("price_id").Optional(),
		field.Int64("plan_id").Optional(),
		field.Uint8("quantity").Default(1),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("orders", Orders.Type).Ref("items").Unique(), // many-to-one
	}
}