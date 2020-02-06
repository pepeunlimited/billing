package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// => `payment_instruments`
type Instrument struct {
	ent.Schema
}

func (Instrument) Config() ent.Config {
	return ent.Config{Table:"payment_instruments"}
}

func (Instrument) Fields() []ent.Field {
	return []ent.Field {
		field.String("type").Unique().MaxLen(12),
		field.Int64("type_i18n_id").Optional(),
	}
}

func (Instrument) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("payments", Payment.Type),
	}
}