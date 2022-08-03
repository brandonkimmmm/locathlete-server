// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"locathlete-server/ent/athlete"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Athlete is the model entity for the Athlete schema.
type Athlete struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Athlete) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case athlete.FieldID:
			values[i] = new(sql.NullInt64)
		case athlete.FieldEmail, athlete.FieldBio, athlete.FieldFirstName, athlete.FieldMiddleName, athlete.FieldLastName:
			values[i] = new(sql.NullString)
		case athlete.FieldCreatedAt, athlete.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Athlete", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Athlete fields.
func (a *Athlete) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case athlete.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case athlete.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case athlete.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				a.Bio = value.String
			}
		case athlete.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				a.FirstName = value.String
			}
		case athlete.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				a.MiddleName = value.String
			}
		case athlete.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				a.LastName = value.String
			}
		case athlete.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case athlete.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Athlete.
// Note that you need to call Athlete.Unwrap() before calling this method if this Athlete
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Athlete) Update() *AthleteUpdateOne {
	return (&AthleteClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Athlete entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Athlete) Unwrap() *Athlete {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Athlete is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Athlete) String() string {
	var builder strings.Builder
	builder.WriteString("Athlete(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(a.Bio)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(a.FirstName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(a.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(a.LastName)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Athletes is a parsable slice of Athlete.
type Athletes []*Athlete

func (a Athletes) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
