package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique().
			Default(func() uuid.UUID { return uuid.Must(uuid.NewV7()) }),
		field.Int64("price"),
		field.Enum("status").
			Values("todo", "in_progress", "done").
			Default("todo"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
	}

}
