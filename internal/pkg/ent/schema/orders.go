package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Orders struct {
	ent.Schema
}

func (Orders) Config() ent.Config {
	return ent.Config{Table:"orders"}
}

func (Orders) Fields() []ent.Field {
	return []ent.Field{}
}

func (Orders) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_txs", OrderTXs.Type),     // one-to-many
		edge.To("order_items", OrderItems.Type), // one-to-many
		edge.To("payment_instruments", PaymentInstruments.Type), // one-to-many
	}
}