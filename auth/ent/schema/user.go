package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ent generates auto-increment int id automatically, so no need to define it here
		field.String("name").
			Unique(),
		field.UUID("uuid", uuid.Must(uuid.NewV7())).
			Unique().
			Default(func() uuid.UUID { return uuid.Must(uuid.NewV7()) }),
	}
}

//func (User) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		// Do not generate an endpoint for DELETE /users/{id}
//		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
//	}
//}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).Unique().Required(),
	}
}
