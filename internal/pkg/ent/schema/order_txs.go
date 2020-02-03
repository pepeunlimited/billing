package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type OrderTXs struct {
	ent.Schema
}

func (OrderTXs) Config() ent.Config {
	return ent.Config{Table:"order_txs"}
}

func (OrderTXs) Fields() []ent.Field {
	return []ent.Field{}
}

func (OrderTXs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Orders.Type).Ref("order_txs").Unique(), // many-to-one
	}
}