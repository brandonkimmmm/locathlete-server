// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"locathlete-server/ent/athlete"
	"locathlete-server/ent/school"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SchoolCreate is the builder for creating a School entity.
type SchoolCreate struct {
	config
	mutation *SchoolMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SchoolCreate) SetName(s string) *SchoolCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SchoolCreate) SetDescription(s string) *SchoolCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetStreetAddress sets the "street_address" field.
func (sc *SchoolCreate) SetStreetAddress(s string) *SchoolCreate {
	sc.mutation.SetStreetAddress(s)
	return sc
}

// SetCity sets the "city" field.
func (sc *SchoolCreate) SetCity(s string) *SchoolCreate {
	sc.mutation.SetCity(s)
	return sc
}

// SetCountry sets the "country" field.
func (sc *SchoolCreate) SetCountry(s string) *SchoolCreate {
	sc.mutation.SetCountry(s)
	return sc
}

// SetAdministrationArea sets the "administration_area" field.
func (sc *SchoolCreate) SetAdministrationArea(s string) *SchoolCreate {
	sc.mutation.SetAdministrationArea(s)
	return sc
}

// SetPostalCode sets the "postal_code" field.
func (sc *SchoolCreate) SetPostalCode(s string) *SchoolCreate {
	sc.mutation.SetPostalCode(s)
	return sc
}

// SetLat sets the "lat" field.
func (sc *SchoolCreate) SetLat(f float64) *SchoolCreate {
	sc.mutation.SetLat(f)
	return sc
}

// SetLng sets the "lng" field.
func (sc *SchoolCreate) SetLng(f float64) *SchoolCreate {
	sc.mutation.SetLng(f)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SchoolCreate) SetCreatedAt(t time.Time) *SchoolCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SchoolCreate) SetNillableCreatedAt(t *time.Time) *SchoolCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SchoolCreate) SetUpdatedAt(t time.Time) *SchoolCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SchoolCreate) SetNillableUpdatedAt(t *time.Time) *SchoolCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// AddAthleteIDs adds the "athletes" edge to the Athlete entity by IDs.
func (sc *SchoolCreate) AddAthleteIDs(ids ...int) *SchoolCreate {
	sc.mutation.AddAthleteIDs(ids...)
	return sc
}

// AddAthletes adds the "athletes" edges to the Athlete entity.
func (sc *SchoolCreate) AddAthletes(a ...*Athlete) *SchoolCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAthleteIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (sc *SchoolCreate) Mutation() *SchoolMutation {
	return sc.mutation
}

// Save creates the School in the database.
func (sc *SchoolCreate) Save(ctx context.Context) (*School, error) {
	var (
		err  error
		node *School
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*School)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SchoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SchoolCreate) SaveX(ctx context.Context) *School {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SchoolCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SchoolCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SchoolCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := school.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := school.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SchoolCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "School.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := school.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "School.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "School.description"`)}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := school.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "School.description": %w`, err)}
		}
	}
	if _, ok := sc.mutation.StreetAddress(); !ok {
		return &ValidationError{Name: "street_address", err: errors.New(`ent: missing required field "School.street_address"`)}
	}
	if v, ok := sc.mutation.StreetAddress(); ok {
		if err := school.StreetAddressValidator(v); err != nil {
			return &ValidationError{Name: "street_address", err: fmt.Errorf(`ent: validator failed for field "School.street_address": %w`, err)}
		}
	}
	if _, ok := sc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "School.city"`)}
	}
	if v, ok := sc.mutation.City(); ok {
		if err := school.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "School.city": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "School.country"`)}
	}
	if v, ok := sc.mutation.Country(); ok {
		if err := school.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "School.country": %w`, err)}
		}
	}
	if _, ok := sc.mutation.AdministrationArea(); !ok {
		return &ValidationError{Name: "administration_area", err: errors.New(`ent: missing required field "School.administration_area"`)}
	}
	if v, ok := sc.mutation.AdministrationArea(); ok {
		if err := school.AdministrationAreaValidator(v); err != nil {
			return &ValidationError{Name: "administration_area", err: fmt.Errorf(`ent: validator failed for field "School.administration_area": %w`, err)}
		}
	}
	if _, ok := sc.mutation.PostalCode(); !ok {
		return &ValidationError{Name: "postal_code", err: errors.New(`ent: missing required field "School.postal_code"`)}
	}
	if v, ok := sc.mutation.PostalCode(); ok {
		if err := school.PostalCodeValidator(v); err != nil {
			return &ValidationError{Name: "postal_code", err: fmt.Errorf(`ent: validator failed for field "School.postal_code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "School.lat"`)}
	}
	if v, ok := sc.mutation.Lat(); ok {
		if err := school.LatValidator(v); err != nil {
			return &ValidationError{Name: "lat", err: fmt.Errorf(`ent: validator failed for field "School.lat": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Lng(); !ok {
		return &ValidationError{Name: "lng", err: errors.New(`ent: missing required field "School.lng"`)}
	}
	if v, ok := sc.mutation.Lng(); ok {
		if err := school.LngValidator(v); err != nil {
			return &ValidationError{Name: "lng", err: fmt.Errorf(`ent: validator failed for field "School.lng": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "School.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "School.updated_at"`)}
	}
	return nil
}

func (sc *SchoolCreate) sqlSave(ctx context.Context) (*School, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SchoolCreate) createSpec() (*School, *sqlgraph.CreateSpec) {
	var (
		_node = &School{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: school.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: school.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := sc.mutation.StreetAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldStreetAddress,
		})
		_node.StreetAddress = value
	}
	if value, ok := sc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldCity,
		})
		_node.City = value
	}
	if value, ok := sc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := sc.mutation.AdministrationArea(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldAdministrationArea,
		})
		_node.AdministrationArea = value
	}
	if value, ok := sc.mutation.PostalCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldPostalCode,
		})
		_node.PostalCode = value
	}
	if value, ok := sc.mutation.Lat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: school.FieldLat,
		})
		_node.Lat = value
	}
	if value, ok := sc.mutation.Lng(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: school.FieldLng,
		})
		_node.Lng = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.AthletesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   school.AthletesTable,
			Columns: school.AthletesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: athlete.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AthleteSchoolCreate{config: sc.config, mutation: newAthleteSchoolMutation(sc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SchoolCreateBulk is the builder for creating many School entities in bulk.
type SchoolCreateBulk struct {
	config
	builders []*SchoolCreate
}

// Save creates the School entities in the database.
func (scb *SchoolCreateBulk) Save(ctx context.Context) ([]*School, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*School, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchoolMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SchoolCreateBulk) SaveX(ctx context.Context) []*School {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SchoolCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SchoolCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
