package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TaskLog holds the schema definition for the TaskLog entity.
type TaskLog struct {
	ent.Schema
}

// Fields of the TaskLog.
func (TaskLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("idempotency_key").Unique(),
	}
}

// Edges of the TaskLog.
func (TaskLog) Edges() []ent.Edge {
	return nil
}
