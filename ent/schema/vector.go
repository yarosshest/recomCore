package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vector holds the schema definition for the Vector entity.
type Vector struct {
	ent.Schema
}

// Fields of the Vector.
func (Vector) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("vector"),
	}
}

// Edges of the Vector.
func (Vector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Product.Type).
			Ref("vectors").
			Unique(),
	}
}
