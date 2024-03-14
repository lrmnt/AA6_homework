// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/analytics/ent/userlog"
)

// UserLogCreate is the builder for creating a UserLog entity.
type UserLogCreate struct {
	config
	mutation *UserLogMutation
	hooks    []Hook
}

// SetIdempotencyKey sets the "idempotency_key" field.
func (ulc *UserLogCreate) SetIdempotencyKey(s string) *UserLogCreate {
	ulc.mutation.SetIdempotencyKey(s)
	return ulc
}

// Mutation returns the UserLogMutation object of the builder.
func (ulc *UserLogCreate) Mutation() *UserLogMutation {
	return ulc.mutation
}

// Save creates the UserLog in the database.
func (ulc *UserLogCreate) Save(ctx context.Context) (*UserLog, error) {
	return withHooks(ctx, ulc.sqlSave, ulc.mutation, ulc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ulc *UserLogCreate) SaveX(ctx context.Context) *UserLog {
	v, err := ulc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulc *UserLogCreate) Exec(ctx context.Context) error {
	_, err := ulc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulc *UserLogCreate) ExecX(ctx context.Context) {
	if err := ulc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulc *UserLogCreate) check() error {
	if _, ok := ulc.mutation.IdempotencyKey(); !ok {
		return &ValidationError{Name: "idempotency_key", err: errors.New(`ent: missing required field "UserLog.idempotency_key"`)}
	}
	return nil
}

func (ulc *UserLogCreate) sqlSave(ctx context.Context) (*UserLog, error) {
	if err := ulc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ulc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ulc.mutation.id = &_node.ID
	ulc.mutation.done = true
	return _node, nil
}

func (ulc *UserLogCreate) createSpec() (*UserLog, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLog{config: ulc.config}
		_spec = sqlgraph.NewCreateSpec(userlog.Table, sqlgraph.NewFieldSpec(userlog.FieldID, field.TypeInt))
	)
	if value, ok := ulc.mutation.IdempotencyKey(); ok {
		_spec.SetField(userlog.FieldIdempotencyKey, field.TypeString, value)
		_node.IdempotencyKey = value
	}
	return _node, _spec
}

// UserLogCreateBulk is the builder for creating many UserLog entities in bulk.
type UserLogCreateBulk struct {
	config
	err      error
	builders []*UserLogCreate
}

// Save creates the UserLog entities in the database.
func (ulcb *UserLogCreateBulk) Save(ctx context.Context) ([]*UserLog, error) {
	if ulcb.err != nil {
		return nil, ulcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ulcb.builders))
	nodes := make([]*UserLog, len(ulcb.builders))
	mutators := make([]Mutator, len(ulcb.builders))
	for i := range ulcb.builders {
		func(i int, root context.Context) {
			builder := ulcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ulcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ulcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulcb *UserLogCreateBulk) SaveX(ctx context.Context) []*UserLog {
	v, err := ulcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulcb *UserLogCreateBulk) Exec(ctx context.Context) error {
	_, err := ulcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulcb *UserLogCreateBulk) ExecX(ctx context.Context) {
	if err := ulcb.Exec(ctx); err != nil {
		panic(err)
	}
}
