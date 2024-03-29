// Code generated by ent, DO NOT EDIT.

package billingcycle

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the billingcycle type in the database.
	Label = "billing_cycle"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldTsFrom holds the string denoting the ts_from field in the database.
	FieldTsFrom = "ts_from"
	// FieldTsTo holds the string denoting the ts_to field in the database.
	FieldTsTo = "ts_to"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// Table holds the table name of the billingcycle in the database.
	Table = "billing_cycles"
)

// Columns holds all SQL columns for billingcycle fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldTsFrom,
	FieldTsTo,
	FieldState,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateOpen   State = "open"
	StateClosed State = "closed"
	StateFuture State = "future"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOpen, StateClosed, StateFuture:
		return nil
	default:
		return fmt.Errorf("billingcycle: invalid enum value for state field: %q", s)
	}
}

// OrderOption defines the ordering options for the BillingCycle queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByTsFrom orders the results by the ts_from field.
func ByTsFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTsFrom, opts...).ToFunc()
}

// ByTsTo orders the results by the ts_to field.
func ByTsTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTsTo, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}
