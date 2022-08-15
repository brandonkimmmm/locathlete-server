package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AthleteSchool holds the schema definition for the AthleteSchool entity.
type AthleteSchool struct {
	ent.Schema
}

func (AthleteSchool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("athlete_id", "school_id"),
	}
}

// Fields of the AthleteSchool.
func (AthleteSchool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("athlete_id"),
		field.Int("school_id"),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the AthleteSchool.
func (AthleteSchool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("athlete", Athlete.Type).Required().Unique().Field("athlete_id"),
		edge.To("school", School.Type).Required().Unique().Field("school_id"),
	}
}
