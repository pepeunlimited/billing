package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type PaymentInstrument struct {
	ent.Schema
}

func (PaymentInstrument) Config() ent.Config {
	return ent.Config{Table:"payment_instruments"}
}

func (PaymentInstrument) Fields() []ent.Field {
	return []ent.Field {
		field.String("a"),
	}
}

func (PaymentInstrument) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("orders", Order.Type),  		// one-to-many
	}
}