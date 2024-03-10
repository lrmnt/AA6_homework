package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserLog holds the schema definition for the UserLog entity.
type UserLog struct {
	ent.Schema
}

// Fields of the UserLog.
func (UserLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("idempotency_key").Unique(),
	}
}

// Edges of the UserLog.
func (UserLog) Edges() []ent.Edge {
	return nil
}
