package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.String("street_address").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("country").NotEmpty(),
		field.String("administration_area").NotEmpty(),
		field.String("postal_code").NotEmpty(),
		field.Float("lat").Range(-90, 90),
		field.Float("lng").Range(-180, 180),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("athletes", Athlete.Type).
			Ref("schools").
			Through("school_athletes", AthleteSchool.Type),
	}
}
