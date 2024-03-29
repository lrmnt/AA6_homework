// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/billing/ent/billingcycle"
	"github.com/lrmnt/AA6_homework/billing/ent/predicate"
)

// BillingCycleDelete is the builder for deleting a BillingCycle entity.
type BillingCycleDelete struct {
	config
	hooks    []Hook
	mutation *BillingCycleMutation
}

// Where appends a list predicates to the BillingCycleDelete builder.
func (bcd *BillingCycleDelete) Where(ps ...predicate.BillingCycle) *BillingCycleDelete {
	bcd.mutation.Where(ps...)
	return bcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcd *BillingCycleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bcd.sqlExec, bcd.mutation, bcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bcd *BillingCycleDelete) ExecX(ctx context.Context) int {
	n, err := bcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcd *BillingCycleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billingcycle.Table, sqlgraph.NewFieldSpec(billingcycle.FieldID, field.TypeInt))
	if ps := bcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bcd.mutation.done = true
	return affected, err
}

// BillingCycleDeleteOne is the builder for deleting a single BillingCycle entity.
type BillingCycleDeleteOne struct {
	bcd *BillingCycleDelete
}

// Where appends a list predicates to the BillingCycleDelete builder.
func (bcdo *BillingCycleDeleteOne) Where(ps ...predicate.BillingCycle) *BillingCycleDeleteOne {
	bcdo.bcd.mutation.Where(ps...)
	return bcdo
}

// Exec executes the deletion query.
func (bcdo *BillingCycleDeleteOne) Exec(ctx context.Context) error {
	n, err := bcdo.bcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billingcycle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdo *BillingCycleDeleteOne) ExecX(ctx context.Context) {
	if err := bcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
