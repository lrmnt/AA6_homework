package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Billing holds the schema definition for the Billing entity.
type Billing struct {
	ent.Schema
}

// Fields of the Billing.
func (Billing) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique(),
		field.Int64("amount"),
		field.Int64("timestamp"),
		field.Enum("type").Values("payout", "debt_move"),
	}
}

// Edges of the Billing.
func (Billing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
		edge.To("billing_cycle", BillingCycle.Type).Unique().Required(),
	}
}
