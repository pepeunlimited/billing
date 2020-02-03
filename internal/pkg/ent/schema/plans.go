package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Plans struct {
	ent.Schema
}

func (Plans) Config() ent.Config {
	return ent.Config{Table:"plans"}
}

func (Plans) Fields() []ent.Field {
	return []ent.Field{}
}

func (Plans) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subscriptions", Subscriptions.Type).Ref("plans").Unique(), // many-to-one
	}
}
