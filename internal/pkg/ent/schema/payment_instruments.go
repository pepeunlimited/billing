package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type PaymentInstruments struct {
	ent.Schema
}

func (PaymentInstruments) Config() ent.Config {
	return ent.Config{Table:"payment_instruments"}
}

func (PaymentInstruments) Fields() []ent.Field {
	return []ent.Field{}
}

func (PaymentInstruments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Orders.Type).Ref("payment_instruments").Unique(), // many-to-one
	}
}