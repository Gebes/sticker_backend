package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gebes.io/sticker_backend/pkg/router/utils/annotations"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (u User) Annotations() []schema.Annotation {
	return annotations.Generate(u)
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stickers", Sticker.Type).Ref("owner"),
	}
}
