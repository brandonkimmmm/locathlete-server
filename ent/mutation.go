// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"locathlete-server/ent/athlete"
	"locathlete-server/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAthlete = "Athlete"
)

// AthleteMutation represents an operation that mutates the Athlete nodes in the graph.
type AthleteMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	bio           *string
	first_name    *string
	middle_name   *string
	last_name     *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Athlete, error)
	predicates    []predicate.Athlete
}

var _ ent.Mutation = (*AthleteMutation)(nil)

// athleteOption allows management of the mutation configuration using functional options.
type athleteOption func(*AthleteMutation)

// newAthleteMutation creates new mutation for the Athlete entity.
func newAthleteMutation(c config, op Op, opts ...athleteOption) *AthleteMutation {
	m := &AthleteMutation{
		config:        c,
		op:            op,
		typ:           TypeAthlete,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAthleteID sets the ID field of the mutation.
func withAthleteID(id int) athleteOption {
	return func(m *AthleteMutation) {
		var (
			err   error
			once  sync.Once
			value *Athlete
		)
		m.oldValue = func(ctx context.Context) (*Athlete, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Athlete.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAthlete sets the old Athlete of the mutation.
func withAthlete(node *Athlete) athleteOption {
	return func(m *AthleteMutation) {
		m.oldValue = func(context.Context) (*Athlete, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AthleteMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AthleteMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AthleteMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AthleteMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Athlete.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *AthleteMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *AthleteMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *AthleteMutation) ResetEmail() {
	m.email = nil
}

// SetBio sets the "bio" field.
func (m *AthleteMutation) SetBio(s string) {
	m.bio = &s
}

// Bio returns the value of the "bio" field in the mutation.
func (m *AthleteMutation) Bio() (r string, exists bool) {
	v := m.bio
	if v == nil {
		return
	}
	return *v, true
}

// OldBio returns the old "bio" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldBio(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBio is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBio requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBio: %w", err)
	}
	return oldValue.Bio, nil
}

// ResetBio resets all changes to the "bio" field.
func (m *AthleteMutation) ResetBio() {
	m.bio = nil
}

// SetFirstName sets the "first_name" field.
func (m *AthleteMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *AthleteMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *AthleteMutation) ResetFirstName() {
	m.first_name = nil
}

// SetMiddleName sets the "middle_name" field.
func (m *AthleteMutation) SetMiddleName(s string) {
	m.middle_name = &s
}

// MiddleName returns the value of the "middle_name" field in the mutation.
func (m *AthleteMutation) MiddleName() (r string, exists bool) {
	v := m.middle_name
	if v == nil {
		return
	}
	return *v, true
}

// OldMiddleName returns the old "middle_name" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldMiddleName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMiddleName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMiddleName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMiddleName: %w", err)
	}
	return oldValue.MiddleName, nil
}

// ClearMiddleName clears the value of the "middle_name" field.
func (m *AthleteMutation) ClearMiddleName() {
	m.middle_name = nil
	m.clearedFields[athlete.FieldMiddleName] = struct{}{}
}

// MiddleNameCleared returns if the "middle_name" field was cleared in this mutation.
func (m *AthleteMutation) MiddleNameCleared() bool {
	_, ok := m.clearedFields[athlete.FieldMiddleName]
	return ok
}

// ResetMiddleName resets all changes to the "middle_name" field.
func (m *AthleteMutation) ResetMiddleName() {
	m.middle_name = nil
	delete(m.clearedFields, athlete.FieldMiddleName)
}

// SetLastName sets the "last_name" field.
func (m *AthleteMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *AthleteMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *AthleteMutation) ResetLastName() {
	m.last_name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *AthleteMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AthleteMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AthleteMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *AthleteMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *AthleteMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Athlete entity.
// If the Athlete object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AthleteMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *AthleteMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the AthleteMutation builder.
func (m *AthleteMutation) Where(ps ...predicate.Athlete) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AthleteMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Athlete).
func (m *AthleteMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AthleteMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.email != nil {
		fields = append(fields, athlete.FieldEmail)
	}
	if m.bio != nil {
		fields = append(fields, athlete.FieldBio)
	}
	if m.first_name != nil {
		fields = append(fields, athlete.FieldFirstName)
	}
	if m.middle_name != nil {
		fields = append(fields, athlete.FieldMiddleName)
	}
	if m.last_name != nil {
		fields = append(fields, athlete.FieldLastName)
	}
	if m.created_at != nil {
		fields = append(fields, athlete.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, athlete.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AthleteMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case athlete.FieldEmail:
		return m.Email()
	case athlete.FieldBio:
		return m.Bio()
	case athlete.FieldFirstName:
		return m.FirstName()
	case athlete.FieldMiddleName:
		return m.MiddleName()
	case athlete.FieldLastName:
		return m.LastName()
	case athlete.FieldCreatedAt:
		return m.CreatedAt()
	case athlete.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AthleteMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case athlete.FieldEmail:
		return m.OldEmail(ctx)
	case athlete.FieldBio:
		return m.OldBio(ctx)
	case athlete.FieldFirstName:
		return m.OldFirstName(ctx)
	case athlete.FieldMiddleName:
		return m.OldMiddleName(ctx)
	case athlete.FieldLastName:
		return m.OldLastName(ctx)
	case athlete.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case athlete.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Athlete field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AthleteMutation) SetField(name string, value ent.Value) error {
	switch name {
	case athlete.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case athlete.FieldBio:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBio(v)
		return nil
	case athlete.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case athlete.FieldMiddleName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMiddleName(v)
		return nil
	case athlete.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case athlete.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case athlete.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Athlete field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AthleteMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AthleteMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AthleteMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Athlete numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AthleteMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(athlete.FieldMiddleName) {
		fields = append(fields, athlete.FieldMiddleName)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AthleteMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AthleteMutation) ClearField(name string) error {
	switch name {
	case athlete.FieldMiddleName:
		m.ClearMiddleName()
		return nil
	}
	return fmt.Errorf("unknown Athlete nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AthleteMutation) ResetField(name string) error {
	switch name {
	case athlete.FieldEmail:
		m.ResetEmail()
		return nil
	case athlete.FieldBio:
		m.ResetBio()
		return nil
	case athlete.FieldFirstName:
		m.ResetFirstName()
		return nil
	case athlete.FieldMiddleName:
		m.ResetMiddleName()
		return nil
	case athlete.FieldLastName:
		m.ResetLastName()
		return nil
	case athlete.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case athlete.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Athlete field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AthleteMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AthleteMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AthleteMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AthleteMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AthleteMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AthleteMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AthleteMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Athlete unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AthleteMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Athlete edge %s", name)
}
