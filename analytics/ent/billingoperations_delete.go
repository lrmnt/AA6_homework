// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/analytics/ent/billingoperations"
	"github.com/lrmnt/AA6_homework/analytics/ent/predicate"
)

// BillingOperationsDelete is the builder for deleting a BillingOperations entity.
type BillingOperationsDelete struct {
	config
	hooks    []Hook
	mutation *BillingOperationsMutation
}

// Where appends a list predicates to the BillingOperationsDelete builder.
func (bod *BillingOperationsDelete) Where(ps ...predicate.BillingOperations) *BillingOperationsDelete {
	bod.mutation.Where(ps...)
	return bod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bod *BillingOperationsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bod.sqlExec, bod.mutation, bod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bod *BillingOperationsDelete) ExecX(ctx context.Context) int {
	n, err := bod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bod *BillingOperationsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billingoperations.Table, sqlgraph.NewFieldSpec(billingoperations.FieldID, field.TypeInt))
	if ps := bod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bod.mutation.done = true
	return affected, err
}

// BillingOperationsDeleteOne is the builder for deleting a single BillingOperations entity.
type BillingOperationsDeleteOne struct {
	bod *BillingOperationsDelete
}

// Where appends a list predicates to the BillingOperationsDelete builder.
func (bodo *BillingOperationsDeleteOne) Where(ps ...predicate.BillingOperations) *BillingOperationsDeleteOne {
	bodo.bod.mutation.Where(ps...)
	return bodo
}

// Exec executes the deletion query.
func (bodo *BillingOperationsDeleteOne) Exec(ctx context.Context) error {
	n, err := bodo.bod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billingoperations.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bodo *BillingOperationsDeleteOne) ExecX(ctx context.Context) {
	if err := bodo.Exec(ctx); err != nil {
		panic(err)
	}
}
