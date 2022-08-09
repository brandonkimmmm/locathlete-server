package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Athlete holds the schema definition for the Athlete entity.
type Athlete struct {
	ent.Schema
}

// Fields of the Athlete.
func (Athlete) Fields() []ent.Field {
	return []ent.Field{
		field.Text("bio").NotEmpty(),
		field.String("first_name").NotEmpty(),
		field.String("middle_name").Optional(),
		field.String("last_name").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Athlete.
func (Athlete) Edges() []ent.Edge {
	return nil
}
