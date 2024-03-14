// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/billingcycle"
	"github.com/lrmnt/AA6_homework/billing/ent/operations"
	"github.com/lrmnt/AA6_homework/billing/ent/task"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
)

// Operations is the model entity for the Operations schema.
type Operations struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int64 `json:"amount,omitempty"`
	// Type holds the value of the "type" field.
	Type operations.Type `json:"type,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OperationsQuery when eager-loading is set.
	Edges                    OperationsEdges `json:"edges"`
	operations_user          *int
	operations_billing_cycle *int
	operations_task          *int
	selectValues             sql.SelectValues
}

// OperationsEdges holds the relations/edges for other nodes in the graph.
type OperationsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// BillingCycle holds the value of the billing_cycle edge.
	BillingCycle *BillingCycle `json:"billing_cycle,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperationsEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BillingCycleOrErr returns the BillingCycle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperationsEdges) BillingCycleOrErr() (*BillingCycle, error) {
	if e.BillingCycle != nil {
		return e.BillingCycle, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: billingcycle.Label}
	}
	return nil, &NotLoadedError{edge: "billing_cycle"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperationsEdges) TaskOrErr() (*Task, error) {
	if e.Task != nil {
		return e.Task, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: task.Label}
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Operations) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case operations.FieldID, operations.FieldAmount:
			values[i] = new(sql.NullInt64)
		case operations.FieldType:
			values[i] = new(sql.NullString)
		case operations.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case operations.FieldUUID:
			values[i] = new(uuid.UUID)
		case operations.ForeignKeys[0]: // operations_user
			values[i] = new(sql.NullInt64)
		case operations.ForeignKeys[1]: // operations_billing_cycle
			values[i] = new(sql.NullInt64)
		case operations.ForeignKeys[2]: // operations_task
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Operations fields.
func (o *Operations) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case operations.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case operations.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				o.UUID = *value
			}
		case operations.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				o.Amount = value.Int64
			}
		case operations.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = operations.Type(value.String)
			}
		case operations.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				o.Timestamp = value.Time
			}
		case operations.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field operations_user", value)
			} else if value.Valid {
				o.operations_user = new(int)
				*o.operations_user = int(value.Int64)
			}
		case operations.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field operations_billing_cycle", value)
			} else if value.Valid {
				o.operations_billing_cycle = new(int)
				*o.operations_billing_cycle = int(value.Int64)
			}
		case operations.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field operations_task", value)
			} else if value.Valid {
				o.operations_task = new(int)
				*o.operations_task = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Operations.
// This includes values selected through modifiers, order, etc.
func (o *Operations) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Operations entity.
func (o *Operations) QueryUser() *UserQuery {
	return NewOperationsClient(o.config).QueryUser(o)
}

// QueryBillingCycle queries the "billing_cycle" edge of the Operations entity.
func (o *Operations) QueryBillingCycle() *BillingCycleQuery {
	return NewOperationsClient(o.config).QueryBillingCycle(o)
}

// QueryTask queries the "task" edge of the Operations entity.
func (o *Operations) QueryTask() *TaskQuery {
	return NewOperationsClient(o.config).QueryTask(o)
}

// Update returns a builder for updating this Operations.
// Note that you need to call Operations.Unwrap() before calling this method if this Operations
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Operations) Update() *OperationsUpdateOne {
	return NewOperationsClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Operations entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Operations) Unwrap() *Operations {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Operations is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Operations) String() string {
	var builder strings.Builder
	builder.WriteString("Operations(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", o.UUID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", o.Amount))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", o.Type))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(o.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OperationsSlice is a parsable slice of Operations.
type OperationsSlice []*Operations