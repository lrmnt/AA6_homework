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
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique(),
		field.String("title"),
		field.String("description"),
		field.Int64("cost"),
		field.Int64("timestamp"),
		field.String("jira_id").Nillable(), // nillable because old tasks stream does not contain jira id
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
	}
}
