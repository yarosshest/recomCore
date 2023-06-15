package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rate holds the schema definition for the Rate entity.
type Rate struct {
	ent.Schema
}

// Fields of the Rate.
func (Rate) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("rate"),
	}
}

// Edges of the Rate.
func (Rate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("rates").
			Unique(),
		edge.From("subject", Product.Type).
			Ref("rates").
			Unique(),
	}
}
