package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Operations holds the schema definition for the Operations entity.
type Operations struct {
	ent.Schema
}

// Fields of the Operations.
func (Operations) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique(),
		field.Int64("amount"),
		field.Enum("type").
			Values("income", "outcome"),
		field.Time("timestamp"),
	}
}

// Edges of the Operations.
func (Operations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
		edge.To("billing_cycle", BillingCycle.Type).Unique().Required(),
		edge.To("task", Task.Type).Unique().Required(),
	}
}
