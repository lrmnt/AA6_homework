// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/billing/ent/predicate"
	"github.com/lrmnt/AA6_homework/billing/ent/userlog"
)

// UserLogUpdate is the builder for updating UserLog entities.
type UserLogUpdate struct {
	config
	hooks    []Hook
	mutation *UserLogMutation
}

// Where appends a list predicates to the UserLogUpdate builder.
func (ulu *UserLogUpdate) Where(ps ...predicate.UserLog) *UserLogUpdate {
	ulu.mutation.Where(ps...)
	return ulu
}

// SetIdempotencyKey sets the "idempotency_key" field.
func (ulu *UserLogUpdate) SetIdempotencyKey(s string) *UserLogUpdate {
	ulu.mutation.SetIdempotencyKey(s)
	return ulu
}

// SetNillableIdempotencyKey sets the "idempotency_key" field if the given value is not nil.
func (ulu *UserLogUpdate) SetNillableIdempotencyKey(s *string) *UserLogUpdate {
	if s != nil {
		ulu.SetIdempotencyKey(*s)
	}
	return ulu
}

// Mutation returns the UserLogMutation object of the builder.
func (ulu *UserLogUpdate) Mutation() *UserLogMutation {
	return ulu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ulu *UserLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ulu.sqlSave, ulu.mutation, ulu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ulu *UserLogUpdate) SaveX(ctx context.Context) int {
	affected, err := ulu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ulu *UserLogUpdate) Exec(ctx context.Context) error {
	_, err := ulu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulu *UserLogUpdate) ExecX(ctx context.Context) {
	if err := ulu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ulu *UserLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userlog.Table, userlog.Columns, sqlgraph.NewFieldSpec(userlog.FieldID, field.TypeInt))
	if ps := ulu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ulu.mutation.IdempotencyKey(); ok {
		_spec.SetField(userlog.FieldIdempotencyKey, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ulu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ulu.mutation.done = true
	return n, nil
}

// UserLogUpdateOne is the builder for updating a single UserLog entity.
type UserLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserLogMutation
}

// SetIdempotencyKey sets the "idempotency_key" field.
func (uluo *UserLogUpdateOne) SetIdempotencyKey(s string) *UserLogUpdateOne {
	uluo.mutation.SetIdempotencyKey(s)
	return uluo
}

// SetNillableIdempotencyKey sets the "idempotency_key" field if the given value is not nil.
func (uluo *UserLogUpdateOne) SetNillableIdempotencyKey(s *string) *UserLogUpdateOne {
	if s != nil {
		uluo.SetIdempotencyKey(*s)
	}
	return uluo
}

// Mutation returns the UserLogMutation object of the builder.
func (uluo *UserLogUpdateOne) Mutation() *UserLogMutation {
	return uluo.mutation
}

// Where appends a list predicates to the UserLogUpdate builder.
func (uluo *UserLogUpdateOne) Where(ps ...predicate.UserLog) *UserLogUpdateOne {
	uluo.mutation.Where(ps...)
	return uluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uluo *UserLogUpdateOne) Select(field string, fields ...string) *UserLogUpdateOne {
	uluo.fields = append([]string{field}, fields...)
	return uluo
}

// Save executes the query and returns the updated UserLog entity.
func (uluo *UserLogUpdateOne) Save(ctx context.Context) (*UserLog, error) {
	return withHooks(ctx, uluo.sqlSave, uluo.mutation, uluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uluo *UserLogUpdateOne) SaveX(ctx context.Context) *UserLog {
	node, err := uluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uluo *UserLogUpdateOne) Exec(ctx context.Context) error {
	_, err := uluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uluo *UserLogUpdateOne) ExecX(ctx context.Context) {
	if err := uluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uluo *UserLogUpdateOne) sqlSave(ctx context.Context) (_node *UserLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(userlog.Table, userlog.Columns, sqlgraph.NewFieldSpec(userlog.FieldID, field.TypeInt))
	id, ok := uluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userlog.FieldID)
		for _, f := range fields {
			if !userlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uluo.mutation.IdempotencyKey(); ok {
		_spec.SetField(userlog.FieldIdempotencyKey, field.TypeString, value)
	}
	_node = &UserLog{config: uluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uluo.mutation.done = true
	return _node, nil
}
