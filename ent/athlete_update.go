// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"locathlete-server/ent/athlete"
	"locathlete-server/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AthleteUpdate is the builder for updating Athlete entities.
type AthleteUpdate struct {
	config
	hooks    []Hook
	mutation *AthleteMutation
}

// Where appends a list predicates to the AthleteUpdate builder.
func (au *AthleteUpdate) Where(ps ...predicate.Athlete) *AthleteUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetEmail sets the "email" field.
func (au *AthleteUpdate) SetEmail(s string) *AthleteUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetBio sets the "bio" field.
func (au *AthleteUpdate) SetBio(s string) *AthleteUpdate {
	au.mutation.SetBio(s)
	return au
}

// SetFirstName sets the "first_name" field.
func (au *AthleteUpdate) SetFirstName(s string) *AthleteUpdate {
	au.mutation.SetFirstName(s)
	return au
}

// SetMiddleName sets the "middle_name" field.
func (au *AthleteUpdate) SetMiddleName(s string) *AthleteUpdate {
	au.mutation.SetMiddleName(s)
	return au
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (au *AthleteUpdate) SetNillableMiddleName(s *string) *AthleteUpdate {
	if s != nil {
		au.SetMiddleName(*s)
	}
	return au
}

// ClearMiddleName clears the value of the "middle_name" field.
func (au *AthleteUpdate) ClearMiddleName() *AthleteUpdate {
	au.mutation.ClearMiddleName()
	return au
}

// SetLastName sets the "last_name" field.
func (au *AthleteUpdate) SetLastName(s string) *AthleteUpdate {
	au.mutation.SetLastName(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AthleteUpdate) SetCreatedAt(t time.Time) *AthleteUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AthleteUpdate) SetNillableCreatedAt(t *time.Time) *AthleteUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AthleteUpdate) SetUpdatedAt(t time.Time) *AthleteUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AthleteMutation object of the builder.
func (au *AthleteUpdate) Mutation() *AthleteMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AthleteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AthleteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AthleteUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AthleteUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AthleteUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AthleteUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := athlete.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AthleteUpdate) check() error {
	if v, ok := au.mutation.Email(); ok {
		if err := athlete.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Athlete.email": %w`, err)}
		}
	}
	if v, ok := au.mutation.FirstName(); ok {
		if err := athlete.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Athlete.first_name": %w`, err)}
		}
	}
	if v, ok := au.mutation.LastName(); ok {
		if err := athlete.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Athlete.last_name": %w`, err)}
		}
	}
	return nil
}

func (au *AthleteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   athlete.Table,
			Columns: athlete.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: athlete.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldEmail,
		})
	}
	if value, ok := au.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldBio,
		})
	}
	if value, ok := au.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldFirstName,
		})
	}
	if value, ok := au.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldMiddleName,
		})
	}
	if au.mutation.MiddleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: athlete.FieldMiddleName,
		})
	}
	if value, ok := au.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldLastName,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: athlete.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: athlete.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{athlete.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AthleteUpdateOne is the builder for updating a single Athlete entity.
type AthleteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AthleteMutation
}

// SetEmail sets the "email" field.
func (auo *AthleteUpdateOne) SetEmail(s string) *AthleteUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetBio sets the "bio" field.
func (auo *AthleteUpdateOne) SetBio(s string) *AthleteUpdateOne {
	auo.mutation.SetBio(s)
	return auo
}

// SetFirstName sets the "first_name" field.
func (auo *AthleteUpdateOne) SetFirstName(s string) *AthleteUpdateOne {
	auo.mutation.SetFirstName(s)
	return auo
}

// SetMiddleName sets the "middle_name" field.
func (auo *AthleteUpdateOne) SetMiddleName(s string) *AthleteUpdateOne {
	auo.mutation.SetMiddleName(s)
	return auo
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (auo *AthleteUpdateOne) SetNillableMiddleName(s *string) *AthleteUpdateOne {
	if s != nil {
		auo.SetMiddleName(*s)
	}
	return auo
}

// ClearMiddleName clears the value of the "middle_name" field.
func (auo *AthleteUpdateOne) ClearMiddleName() *AthleteUpdateOne {
	auo.mutation.ClearMiddleName()
	return auo
}

// SetLastName sets the "last_name" field.
func (auo *AthleteUpdateOne) SetLastName(s string) *AthleteUpdateOne {
	auo.mutation.SetLastName(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AthleteUpdateOne) SetCreatedAt(t time.Time) *AthleteUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AthleteUpdateOne) SetNillableCreatedAt(t *time.Time) *AthleteUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AthleteUpdateOne) SetUpdatedAt(t time.Time) *AthleteUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AthleteMutation object of the builder.
func (auo *AthleteUpdateOne) Mutation() *AthleteMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AthleteUpdateOne) Select(field string, fields ...string) *AthleteUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Athlete entity.
func (auo *AthleteUpdateOne) Save(ctx context.Context) (*Athlete, error) {
	var (
		err  error
		node *Athlete
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AthleteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Athlete)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AthleteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AthleteUpdateOne) SaveX(ctx context.Context) *Athlete {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AthleteUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AthleteUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AthleteUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := athlete.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AthleteUpdateOne) check() error {
	if v, ok := auo.mutation.Email(); ok {
		if err := athlete.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Athlete.email": %w`, err)}
		}
	}
	if v, ok := auo.mutation.FirstName(); ok {
		if err := athlete.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Athlete.first_name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.LastName(); ok {
		if err := athlete.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Athlete.last_name": %w`, err)}
		}
	}
	return nil
}

func (auo *AthleteUpdateOne) sqlSave(ctx context.Context) (_node *Athlete, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   athlete.Table,
			Columns: athlete.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: athlete.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Athlete.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, athlete.FieldID)
		for _, f := range fields {
			if !athlete.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != athlete.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldEmail,
		})
	}
	if value, ok := auo.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldBio,
		})
	}
	if value, ok := auo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldFirstName,
		})
	}
	if value, ok := auo.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldMiddleName,
		})
	}
	if auo.mutation.MiddleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: athlete.FieldMiddleName,
		})
	}
	if value, ok := auo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: athlete.FieldLastName,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: athlete.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: athlete.FieldUpdatedAt,
		})
	}
	_node = &Athlete{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{athlete.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}