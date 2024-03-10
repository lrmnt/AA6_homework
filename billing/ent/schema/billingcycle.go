package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BillingCycle holds the schema definition for the BillingCycle entity.
type BillingCycle struct {
	ent.Schema
}

// Fields of the BillingCycle.
func (BillingCycle) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique(),
		field.Int64("ts_from"),
		field.Int64("ts_to"),
		field.Enum("state").
			Values("open", "closed", "future"),
	}
}

// Edges of the BillingCycle.
func (BillingCycle) Edges() []ent.Edge {
	return nil
}
