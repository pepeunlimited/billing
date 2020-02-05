package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Txs struct {
	ent.Schema
}

func (Txs) Config() ent.Config {
	return ent.Config{Table:"txs"}
}

func (Txs) Fields() []ent.Field {
	return []ent.Field {
		field.String("a"),
	}
}

func (Txs) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("orders", Order.Type).Ref("txs").Unique(), // many-to-one
	}
}