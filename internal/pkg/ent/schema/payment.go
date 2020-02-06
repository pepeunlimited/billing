package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// => `payments`
type Payment struct {
	ent.Schema
}

func (Payment) Config() ent.Config {
	return ent.Config{Table:"payments"}
}

func (Payment) Fields() []ent.Field {
	return []ent.Field{

	}
}

func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Orders.Type).Unique(),
		edge.From("instruments", Instrument.Type).Ref("payments").Unique(), // many-to-one
	}
}