package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Item struct {
	ent.Schema
}

func (Item) Config() ent.Config {
	return ent.Config{Table:"items"}
}

func (Item) Fields() []ent.Field {
	return []ent.Field {
		field.String("a"),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("orders", Order.Type).Ref("items").Unique(), // many-to-one
	}
}