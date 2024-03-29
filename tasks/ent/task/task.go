// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tasks"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "task_user"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldUUID,
	FieldPrice,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusTodo is the default value of the Status enum.
const DefaultStatus = StatusTodo

// Status values.
const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
