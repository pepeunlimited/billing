package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Subscriptions struct {
	ent.Schema
}

func (Subscriptions) Config() ent.Config {
	return ent.Config{Table:"subscriptions"}
}

func (Subscriptions) Fields() []ent.Field {
	return []ent.Field{}
}

func (Subscriptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plans", Plans.Type), // one-to-many
	}
}
