package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gebes.io/sticker_backend/pkg/router/utils/annotations"
	"time"
)

// Sticker holds the schema definition for the Sticker entity.
type Sticker struct {
	ent.Schema
}

func (u Sticker) Annotations() []schema.Annotation {
	return annotations.Generate(u)
}

// Fields of the Sticker.
func (Sticker) Fields() []ent.Field {
	return []ent.Field{
		field.String("location_description").NotEmpty(),
		field.Float("latitude"),
		field.Float("longitude"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Sticker.
func (Sticker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).Unique().Required(),
	}
}
