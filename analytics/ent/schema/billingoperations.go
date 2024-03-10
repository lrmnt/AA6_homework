package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BillingOperations holds the schema definition for the BillingOperations entity.
type BillingOperations struct {
	ent.Schema
}

// Fields of the BillingOperations.
func (BillingOperations) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique(),
		field.Int64("amount"),
		field.Int64("timestamp"),
		field.Enum("type").Values("payout", "debt_move"),
	}
}

// Edges of the BillingOperations.
func (BillingOperations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
	}
}
