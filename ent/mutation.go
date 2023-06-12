// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chris-han-nih/ent-demo/ent/predicate"
	"github.com/chris-han-nih/ent-demo/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	rank          *float64
	addrank       *float64
	active        *bool
	name          *string
	created_at    *time.Time
	url           **url.URL
	strings       *[]string
	appendstrings []string
	state         *user.State
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetRank sets the "rank" field.
func (m *UserMutation) SetRank(f float64) {
	m.rank = &f
	m.addrank = nil
}

// Rank returns the value of the "rank" field in the mutation.
func (m *UserMutation) Rank() (r float64, exists bool) {
	v := m.rank
	if v == nil {
		return
	}
	return *v, true
}

// OldRank returns the old "rank" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRank(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRank is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRank requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRank: %w", err)
	}
	return oldValue.Rank, nil
}

// AddRank adds f to the "rank" field.
func (m *UserMutation) AddRank(f float64) {
	if m.addrank != nil {
		*m.addrank += f
	} else {
		m.addrank = &f
	}
}

// AddedRank returns the value that was added to the "rank" field in this mutation.
func (m *UserMutation) AddedRank() (r float64, exists bool) {
	v := m.addrank
	if v == nil {
		return
	}
	return *v, true
}

// ClearRank clears the value of the "rank" field.
func (m *UserMutation) ClearRank() {
	m.rank = nil
	m.addrank = nil
	m.clearedFields[user.FieldRank] = struct{}{}
}

// RankCleared returns if the "rank" field was cleared in this mutation.
func (m *UserMutation) RankCleared() bool {
	_, ok := m.clearedFields[user.FieldRank]
	return ok
}

// ResetRank resets all changes to the "rank" field.
func (m *UserMutation) ResetRank() {
	m.rank = nil
	m.addrank = nil
	delete(m.clearedFields, user.FieldRank)
}

// SetActive sets the "active" field.
func (m *UserMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the value of the "active" field in the mutation.
func (m *UserMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// OldActive returns the old "active" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActive: %w", err)
	}
	return oldValue.Active, nil
}

// ResetActive resets all changes to the "active" field.
func (m *UserMutation) ResetActive() {
	m.active = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetURL sets the "url" field.
func (m *UserMutation) SetURL(u *url.URL) {
	m.url = &u
}

// URL returns the value of the "url" field in the mutation.
func (m *UserMutation) URL() (r *url.URL, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldURL(ctx context.Context) (v *url.URL, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ClearURL clears the value of the "url" field.
func (m *UserMutation) ClearURL() {
	m.url = nil
	m.clearedFields[user.FieldURL] = struct{}{}
}

// URLCleared returns if the "url" field was cleared in this mutation.
func (m *UserMutation) URLCleared() bool {
	_, ok := m.clearedFields[user.FieldURL]
	return ok
}

// ResetURL resets all changes to the "url" field.
func (m *UserMutation) ResetURL() {
	m.url = nil
	delete(m.clearedFields, user.FieldURL)
}

// SetStrings sets the "strings" field.
func (m *UserMutation) SetStrings(s []string) {
	m.strings = &s
	m.appendstrings = nil
}

// Strings returns the value of the "strings" field in the mutation.
func (m *UserMutation) Strings() (r []string, exists bool) {
	v := m.strings
	if v == nil {
		return
	}
	return *v, true
}

// OldStrings returns the old "strings" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStrings(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStrings is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStrings requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStrings: %w", err)
	}
	return oldValue.Strings, nil
}

// AppendStrings adds s to the "strings" field.
func (m *UserMutation) AppendStrings(s []string) {
	m.appendstrings = append(m.appendstrings, s...)
}

// AppendedStrings returns the list of values that were appended to the "strings" field in this mutation.
func (m *UserMutation) AppendedStrings() ([]string, bool) {
	if len(m.appendstrings) == 0 {
		return nil, false
	}
	return m.appendstrings, true
}

// ClearStrings clears the value of the "strings" field.
func (m *UserMutation) ClearStrings() {
	m.strings = nil
	m.appendstrings = nil
	m.clearedFields[user.FieldStrings] = struct{}{}
}

// StringsCleared returns if the "strings" field was cleared in this mutation.
func (m *UserMutation) StringsCleared() bool {
	_, ok := m.clearedFields[user.FieldStrings]
	return ok
}

// ResetStrings resets all changes to the "strings" field.
func (m *UserMutation) ResetStrings() {
	m.strings = nil
	m.appendstrings = nil
	delete(m.clearedFields, user.FieldStrings)
}

// SetState sets the "state" field.
func (m *UserMutation) SetState(u user.State) {
	m.state = &u
}

// State returns the value of the "state" field in the mutation.
func (m *UserMutation) State() (r user.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldState(ctx context.Context) (v user.State, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ClearState clears the value of the "state" field.
func (m *UserMutation) ClearState() {
	m.state = nil
	m.clearedFields[user.FieldState] = struct{}{}
}

// StateCleared returns if the "state" field was cleared in this mutation.
func (m *UserMutation) StateCleared() bool {
	_, ok := m.clearedFields[user.FieldState]
	return ok
}

// ResetState resets all changes to the "state" field.
func (m *UserMutation) ResetState() {
	m.state = nil
	delete(m.clearedFields, user.FieldState)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.rank != nil {
		fields = append(fields, user.FieldRank)
	}
	if m.active != nil {
		fields = append(fields, user.FieldActive)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.url != nil {
		fields = append(fields, user.FieldURL)
	}
	if m.strings != nil {
		fields = append(fields, user.FieldStrings)
	}
	if m.state != nil {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldRank:
		return m.Rank()
	case user.FieldActive:
		return m.Active()
	case user.FieldName:
		return m.Name()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldURL:
		return m.URL()
	case user.FieldStrings:
		return m.Strings()
	case user.FieldState:
		return m.State()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldRank:
		return m.OldRank(ctx)
	case user.FieldActive:
		return m.OldActive(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldURL:
		return m.OldURL(ctx)
	case user.FieldStrings:
		return m.OldStrings(ctx)
	case user.FieldState:
		return m.OldState(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldRank:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRank(v)
		return nil
	case user.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldURL:
		v, ok := value.(*url.URL)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case user.FieldStrings:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStrings(v)
		return nil
	case user.FieldState:
		v, ok := value.(user.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.addrank != nil {
		fields = append(fields, user.FieldRank)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	case user.FieldRank:
		return m.AddedRank()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	case user.FieldRank:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRank(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldRank) {
		fields = append(fields, user.FieldRank)
	}
	if m.FieldCleared(user.FieldURL) {
		fields = append(fields, user.FieldURL)
	}
	if m.FieldCleared(user.FieldStrings) {
		fields = append(fields, user.FieldStrings)
	}
	if m.FieldCleared(user.FieldState) {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldRank:
		m.ClearRank()
		return nil
	case user.FieldURL:
		m.ClearURL()
		return nil
	case user.FieldStrings:
		m.ClearStrings()
		return nil
	case user.FieldState:
		m.ClearState()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldRank:
		m.ResetRank()
		return nil
	case user.FieldActive:
		m.ResetActive()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldURL:
		m.ResetURL()
		return nil
	case user.FieldStrings:
		m.ResetStrings()
		return nil
	case user.FieldState:
		m.ResetState()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
