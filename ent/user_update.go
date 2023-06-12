// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/chris-han-nih/ent-demo/ent/predicate"
	"github.com/chris-han-nih/ent-demo/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetRank sets the "rank" field.
func (uu *UserUpdate) SetRank(f float64) *UserUpdate {
	uu.mutation.ResetRank()
	uu.mutation.SetRank(f)
	return uu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRank(f *float64) *UserUpdate {
	if f != nil {
		uu.SetRank(*f)
	}
	return uu
}

// AddRank adds f to the "rank" field.
func (uu *UserUpdate) AddRank(f float64) *UserUpdate {
	uu.mutation.AddRank(f)
	return uu
}

// ClearRank clears the value of the "rank" field.
func (uu *UserUpdate) ClearRank() *UserUpdate {
	uu.mutation.ClearRank()
	return uu
}

// SetActive sets the "active" field.
func (uu *UserUpdate) SetActive(b bool) *UserUpdate {
	uu.mutation.SetActive(b)
	return uu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetActive(*b)
	}
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetURL sets the "url" field.
func (uu *UserUpdate) SetURL(u *url.URL) *UserUpdate {
	uu.mutation.SetURL(u)
	return uu
}

// ClearURL clears the value of the "url" field.
func (uu *UserUpdate) ClearURL() *UserUpdate {
	uu.mutation.ClearURL()
	return uu
}

// SetStrings sets the "strings" field.
func (uu *UserUpdate) SetStrings(s []string) *UserUpdate {
	uu.mutation.SetStrings(s)
	return uu
}

// AppendStrings appends s to the "strings" field.
func (uu *UserUpdate) AppendStrings(s []string) *UserUpdate {
	uu.mutation.AppendStrings(s)
	return uu
}

// ClearStrings clears the value of the "strings" field.
func (uu *UserUpdate) ClearStrings() *UserUpdate {
	uu.mutation.ClearStrings()
	return uu
}

// SetState sets the "state" field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.mutation.SetState(u)
	return uu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of the "state" field.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.mutation.ClearState()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	if v, ok := uu.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "User.state": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeFloat64, value)
	}
	if uu.mutation.RankCleared() {
		_spec.ClearField(user.FieldRank, field.TypeFloat64)
	}
	if value, ok := uu.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uu.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uu.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uu.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uu.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if uu.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetRank sets the "rank" field.
func (uuo *UserUpdateOne) SetRank(f float64) *UserUpdateOne {
	uuo.mutation.ResetRank()
	uuo.mutation.SetRank(f)
	return uuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRank(f *float64) *UserUpdateOne {
	if f != nil {
		uuo.SetRank(*f)
	}
	return uuo
}

// AddRank adds f to the "rank" field.
func (uuo *UserUpdateOne) AddRank(f float64) *UserUpdateOne {
	uuo.mutation.AddRank(f)
	return uuo
}

// ClearRank clears the value of the "rank" field.
func (uuo *UserUpdateOne) ClearRank() *UserUpdateOne {
	uuo.mutation.ClearRank()
	return uuo
}

// SetActive sets the "active" field.
func (uuo *UserUpdateOne) SetActive(b bool) *UserUpdateOne {
	uuo.mutation.SetActive(b)
	return uuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetActive(*b)
	}
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetURL sets the "url" field.
func (uuo *UserUpdateOne) SetURL(u *url.URL) *UserUpdateOne {
	uuo.mutation.SetURL(u)
	return uuo
}

// ClearURL clears the value of the "url" field.
func (uuo *UserUpdateOne) ClearURL() *UserUpdateOne {
	uuo.mutation.ClearURL()
	return uuo
}

// SetStrings sets the "strings" field.
func (uuo *UserUpdateOne) SetStrings(s []string) *UserUpdateOne {
	uuo.mutation.SetStrings(s)
	return uuo
}

// AppendStrings appends s to the "strings" field.
func (uuo *UserUpdateOne) AppendStrings(s []string) *UserUpdateOne {
	uuo.mutation.AppendStrings(s)
	return uuo
}

// ClearStrings clears the value of the "strings" field.
func (uuo *UserUpdateOne) ClearStrings() *UserUpdateOne {
	uuo.mutation.ClearStrings()
	return uuo
}

// SetState sets the "state" field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.mutation.SetState(u)
	return uuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of the "state" field.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.mutation.ClearState()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "User.state": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeFloat64, value)
	}
	if uuo.mutation.RankCleared() {
		_spec.ClearField(user.FieldRank, field.TypeFloat64)
	}
	if value, ok := uuo.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uuo.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uuo.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uuo.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if uuo.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
