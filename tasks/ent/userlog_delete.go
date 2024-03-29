// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/tasks/ent/predicate"
	"github.com/lrmnt/AA6_homework/tasks/ent/userlog"
)

// UserLogDelete is the builder for deleting a UserLog entity.
type UserLogDelete struct {
	config
	hooks    []Hook
	mutation *UserLogMutation
}

// Where appends a list predicates to the UserLogDelete builder.
func (uld *UserLogDelete) Where(ps ...predicate.UserLog) *UserLogDelete {
	uld.mutation.Where(ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UserLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uld.sqlExec, uld.mutation, uld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UserLogDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UserLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userlog.Table, sqlgraph.NewFieldSpec(userlog.FieldID, field.TypeInt))
	if ps := uld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uld.mutation.done = true
	return affected, err
}

// UserLogDeleteOne is the builder for deleting a single UserLog entity.
type UserLogDeleteOne struct {
	uld *UserLogDelete
}

// Where appends a list predicates to the UserLogDelete builder.
func (uldo *UserLogDeleteOne) Where(ps ...predicate.UserLog) *UserLogDeleteOne {
	uldo.uld.mutation.Where(ps...)
	return uldo
}

// Exec executes the deletion query.
func (uldo *UserLogDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UserLogDeleteOne) ExecX(ctx context.Context) {
	if err := uldo.Exec(ctx); err != nil {
		panic(err)
	}
}
