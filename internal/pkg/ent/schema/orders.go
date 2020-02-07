package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"

)

// => `orders`
type Orders struct {
	ent.Schema
}

func (Orders) Config() ent.Config {
	return ent.Config{Table:"orders"}
}

func (Orders) Fields() []ent.Field {
	return []ent.Field {
		field.Time("created_at"),
		field.Int64("user_id"),
	}
}

func (Orders) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("txs", 	  Txs.Type),    							// one-to-many
		edge.To("items", 	  Item.Type), 								// one-to-many
		edge.To("payments", Payment.Type).Unique(), 						// one-to-one
	}
}