// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/task"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tc *TaskCreate) SetUUID(u uuid.UUID) *TaskCreate {
	tc.mutation.SetUUID(u)
	return tc
}

// SetTitle sets the "title" field.
func (tc *TaskCreate) SetTitle(s string) *TaskCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TaskCreate) SetDescription(s string) *TaskCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetCost sets the "cost" field.
func (tc *TaskCreate) SetCost(i int64) *TaskCreate {
	tc.mutation.SetCost(i)
	return tc
}

// SetJiraID sets the "jira_id" field.
func (tc *TaskCreate) SetJiraID(s string) *TaskCreate {
	tc.mutation.SetJiraID(s)
	return tc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tc *TaskCreate) SetUserID(id int) *TaskCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TaskCreate) SetUser(u *User) *TaskCreate {
	return tc.SetUserID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Task.uuid"`)}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Task.title"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Task.description"`)}
	}
	if _, ok := tc.mutation.Cost(); !ok {
		return &ValidationError{Name: "cost", err: errors.New(`ent: missing required field "Task.cost"`)}
	}
	if _, ok := tc.mutation.JiraID(); !ok {
		return &ValidationError{Name: "jira_id", err: errors.New(`ent: missing required field "Task.jira_id"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Task.user"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Cost(); ok {
		_spec.SetField(task.FieldCost, field.TypeInt64, value)
		_node.Cost = value
	}
	if value, ok := tc.mutation.JiraID(); ok {
		_spec.SetField(task.FieldJiraID, field.TypeString, value)
		_node.JiraID = &value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}