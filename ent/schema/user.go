package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	LegalAge = 18
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("age").Positive().Min(LegalAge),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
