package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value_type").Optional(),
		field.String("value"),
		field.String("value_description").Optional(),
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Product.Type).
			Ref("attributes").
			Unique(),
	}
}
