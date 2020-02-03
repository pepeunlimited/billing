package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type OrderItems struct {
	ent.Schema
}

func (OrderItems) Config() ent.Config {
	return ent.Config{Table:"order_items"}
}

func (OrderItems) Fields() []ent.Field {
	return []ent.Field{}
}

func (OrderItems) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Orders.Type).Ref("order_items").Unique(), // many-to-one
	}
}