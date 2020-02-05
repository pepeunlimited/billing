package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Config() ent.Config {
	return ent.Config{Table:"orders"}
}

func (Order) Fields() []ent.Field {
	return []ent.Field {
		field.String("a"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("txs", Txs.Type),               	 												// one-to-many
		edge.To("items", Item.Type),                  											// one-to-many
		edge.From("payment_instruments", PaymentInstrument.Type).Ref("orders").Unique(), 		// one-to-many
	}
}