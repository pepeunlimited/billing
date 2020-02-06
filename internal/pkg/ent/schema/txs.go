package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// => `order_txs`
type Txs struct {
	ent.Schema
}

func (Txs) Config() ent.Config {
	return ent.Config{Table:"order_txs"}
}

func (Txs) Fields() []ent.Field {
	return []ent.Field {
		field.String("status").MaxLen(8),
		field.Time("created_at"),
	}
}

func (Txs) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("orders", Orders.Type).Ref("txs").Unique(), // many-to-one
	}
}