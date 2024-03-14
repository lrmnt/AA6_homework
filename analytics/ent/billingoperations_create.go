// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/analytics/ent/billingoperations"
	"github.com/lrmnt/AA6_homework/analytics/ent/user"
)

// BillingOperationsCreate is the builder for creating a BillingOperations entity.
type BillingOperationsCreate struct {
	config
	mutation *BillingOperationsMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (boc *BillingOperationsCreate) SetUUID(u uuid.UUID) *BillingOperationsCreate {
	boc.mutation.SetUUID(u)
	return boc
}

// SetAmount sets the "amount" field.
func (boc *BillingOperationsCreate) SetAmount(i int64) *BillingOperationsCreate {
	boc.mutation.SetAmount(i)
	return boc
}

// SetTimestamp sets the "timestamp" field.
func (boc *BillingOperationsCreate) SetTimestamp(i int64) *BillingOperationsCreate {
	boc.mutation.SetTimestamp(i)
	return boc
}

// SetType sets the "type" field.
func (boc *BillingOperationsCreate) SetType(b billingoperations.Type) *BillingOperationsCreate {
	boc.mutation.SetType(b)
	return boc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (boc *BillingOperationsCreate) SetUserID(id int) *BillingOperationsCreate {
	boc.mutation.SetUserID(id)
	return boc
}

// SetUser sets the "user" edge to the User entity.
func (boc *BillingOperationsCreate) SetUser(u *User) *BillingOperationsCreate {
	return boc.SetUserID(u.ID)
}

// Mutation returns the BillingOperationsMutation object of the builder.
func (boc *BillingOperationsCreate) Mutation() *BillingOperationsMutation {
	return boc.mutation
}

// Save creates the BillingOperations in the database.
func (boc *BillingOperationsCreate) Save(ctx context.Context) (*BillingOperations, error) {
	return withHooks(ctx, boc.sqlSave, boc.mutation, boc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (boc *BillingOperationsCreate) SaveX(ctx context.Context) *BillingOperations {
	v, err := boc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (boc *BillingOperationsCreate) Exec(ctx context.Context) error {
	_, err := boc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (boc *BillingOperationsCreate) ExecX(ctx context.Context) {
	if err := boc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (boc *BillingOperationsCreate) check() error {
	if _, ok := boc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "BillingOperations.uuid"`)}
	}
	if _, ok := boc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "BillingOperations.amount"`)}
	}
	if _, ok := boc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "BillingOperations.timestamp"`)}
	}
	if _, ok := boc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "BillingOperations.type"`)}
	}
	if v, ok := boc.mutation.GetType(); ok {
		if err := billingoperations.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "BillingOperations.type": %w`, err)}
		}
	}
	if _, ok := boc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "BillingOperations.user"`)}
	}
	return nil
}

func (boc *BillingOperationsCreate) sqlSave(ctx context.Context) (*BillingOperations, error) {
	if err := boc.check(); err != nil {
		return nil, err
	}
	_node, _spec := boc.createSpec()
	if err := sqlgraph.CreateNode(ctx, boc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	boc.mutation.id = &_node.ID
	boc.mutation.done = true
	return _node, nil
}

func (boc *BillingOperationsCreate) createSpec() (*BillingOperations, *sqlgraph.CreateSpec) {
	var (
		_node = &BillingOperations{config: boc.config}
		_spec = sqlgraph.NewCreateSpec(billingoperations.Table, sqlgraph.NewFieldSpec(billingoperations.FieldID, field.TypeInt))
	)
	if value, ok := boc.mutation.UUID(); ok {
		_spec.SetField(billingoperations.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := boc.mutation.Amount(); ok {
		_spec.SetField(billingoperations.FieldAmount, field.TypeInt64, value)
		_node.Amount = value
	}
	if value, ok := boc.mutation.Timestamp(); ok {
		_spec.SetField(billingoperations.FieldTimestamp, field.TypeInt64, value)
		_node.Timestamp = value
	}
	if value, ok := boc.mutation.GetType(); ok {
		_spec.SetField(billingoperations.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := boc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billingoperations.UserTable,
			Columns: []string{billingoperations.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.billing_operations_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BillingOperationsCreateBulk is the builder for creating many BillingOperations entities in bulk.
type BillingOperationsCreateBulk struct {
	config
	err      error
	builders []*BillingOperationsCreate
}

// Save creates the BillingOperations entities in the database.
func (bocb *BillingOperationsCreateBulk) Save(ctx context.Context) ([]*BillingOperations, error) {
	if bocb.err != nil {
		return nil, bocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bocb.builders))
	nodes := make([]*BillingOperations, len(bocb.builders))
	mutators := make([]Mutator, len(bocb.builders))
	for i := range bocb.builders {
		func(i int, root context.Context) {
			builder := bocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingOperationsMutation)
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
					_, err = mutators[i+1].Mutate(root, bocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bocb *BillingOperationsCreateBulk) SaveX(ctx context.Context) []*BillingOperations {
	v, err := bocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bocb *BillingOperationsCreateBulk) Exec(ctx context.Context) error {
	_, err := bocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bocb *BillingOperationsCreateBulk) ExecX(ctx context.Context) {
	if err := bocb.Exec(ctx); err != nil {
		panic(err)
	}
}