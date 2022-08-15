// Code generated by ent, DO NOT EDIT.

package athleteschool

import (
	"locathlete-server/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// AthleteID applies equality check predicate on the "athlete_id" field. It's identical to AthleteIDEQ.
func AthleteID(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAthleteID), v))
	})
}

// SchoolID applies equality check predicate on the "school_id" field. It's identical to SchoolIDEQ.
func SchoolID(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchoolID), v))
	})
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// AthleteIDEQ applies the EQ predicate on the "athlete_id" field.
func AthleteIDEQ(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAthleteID), v))
	})
}

// AthleteIDNEQ applies the NEQ predicate on the "athlete_id" field.
func AthleteIDNEQ(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAthleteID), v))
	})
}

// AthleteIDIn applies the In predicate on the "athlete_id" field.
func AthleteIDIn(vs ...int) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAthleteID), v...))
	})
}

// AthleteIDNotIn applies the NotIn predicate on the "athlete_id" field.
func AthleteIDNotIn(vs ...int) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAthleteID), v...))
	})
}

// SchoolIDEQ applies the EQ predicate on the "school_id" field.
func SchoolIDEQ(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchoolID), v))
	})
}

// SchoolIDNEQ applies the NEQ predicate on the "school_id" field.
func SchoolIDNEQ(v int) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSchoolID), v))
	})
}

// SchoolIDIn applies the In predicate on the "school_id" field.
func SchoolIDIn(vs ...int) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSchoolID), v...))
	})
}

// SchoolIDNotIn applies the NotIn predicate on the "school_id" field.
func SchoolIDNotIn(vs ...int) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSchoolID), v...))
	})
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndDate), v))
	})
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndDate), v...))
	})
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndDate), v...))
	})
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndDate), v))
	})
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndDate), v))
	})
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndDate), v))
	})
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndDate), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AthleteSchool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AthleteSchool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasAthlete applies the HasEdge predicate on the "athlete" edge.
func HasAthlete() predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, AthleteColumn),
			sqlgraph.To(AthleteInverseTable, AthleteFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AthleteTable, AthleteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAthleteWith applies the HasEdge predicate on the "athlete" edge with a given conditions (other predicates).
func HasAthleteWith(preds ...predicate.Athlete) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, AthleteColumn),
			sqlgraph.To(AthleteInverseTable, AthleteFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AthleteTable, AthleteColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSchool applies the HasEdge predicate on the "school" edge.
func HasSchool() predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, SchoolColumn),
			sqlgraph.To(SchoolInverseTable, SchoolFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SchoolTable, SchoolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchoolWith applies the HasEdge predicate on the "school" edge with a given conditions (other predicates).
func HasSchoolWith(preds ...predicate.School) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, SchoolColumn),
			sqlgraph.To(SchoolInverseTable, SchoolFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SchoolTable, SchoolColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AthleteSchool) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AthleteSchool) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AthleteSchool) predicate.AthleteSchool {
	return predicate.AthleteSchool(func(s *sql.Selector) {
		p(s.Not())
	})
}
