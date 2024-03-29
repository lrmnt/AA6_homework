// Code generated by ent, DO NOT EDIT.

package billingcycle

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldUUID, v))
}

// TsFrom applies equality check predicate on the "ts_from" field. It's identical to TsFromEQ.
func TsFrom(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldTsFrom, v))
}

// TsTo applies equality check predicate on the "ts_to" field. It's identical to TsToEQ.
func TsTo(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldTsTo, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLTE(FieldUUID, v))
}

// TsFromEQ applies the EQ predicate on the "ts_from" field.
func TsFromEQ(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldTsFrom, v))
}

// TsFromNEQ applies the NEQ predicate on the "ts_from" field.
func TsFromNEQ(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNEQ(FieldTsFrom, v))
}

// TsFromIn applies the In predicate on the "ts_from" field.
func TsFromIn(vs ...int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldIn(FieldTsFrom, vs...))
}

// TsFromNotIn applies the NotIn predicate on the "ts_from" field.
func TsFromNotIn(vs ...int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNotIn(FieldTsFrom, vs...))
}

// TsFromGT applies the GT predicate on the "ts_from" field.
func TsFromGT(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGT(FieldTsFrom, v))
}

// TsFromGTE applies the GTE predicate on the "ts_from" field.
func TsFromGTE(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGTE(FieldTsFrom, v))
}

// TsFromLT applies the LT predicate on the "ts_from" field.
func TsFromLT(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLT(FieldTsFrom, v))
}

// TsFromLTE applies the LTE predicate on the "ts_from" field.
func TsFromLTE(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLTE(FieldTsFrom, v))
}

// TsToEQ applies the EQ predicate on the "ts_to" field.
func TsToEQ(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldTsTo, v))
}

// TsToNEQ applies the NEQ predicate on the "ts_to" field.
func TsToNEQ(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNEQ(FieldTsTo, v))
}

// TsToIn applies the In predicate on the "ts_to" field.
func TsToIn(vs ...int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldIn(FieldTsTo, vs...))
}

// TsToNotIn applies the NotIn predicate on the "ts_to" field.
func TsToNotIn(vs ...int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNotIn(FieldTsTo, vs...))
}

// TsToGT applies the GT predicate on the "ts_to" field.
func TsToGT(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGT(FieldTsTo, v))
}

// TsToGTE applies the GTE predicate on the "ts_to" field.
func TsToGTE(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldGTE(FieldTsTo, v))
}

// TsToLT applies the LT predicate on the "ts_to" field.
func TsToLT(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLT(FieldTsTo, v))
}

// TsToLTE applies the LTE predicate on the "ts_to" field.
func TsToLTE(v int64) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldLTE(FieldTsTo, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.BillingCycle {
	return predicate.BillingCycle(sql.FieldNotIn(FieldState, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingCycle) predicate.BillingCycle {
	return predicate.BillingCycle(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingCycle) predicate.BillingCycle {
	return predicate.BillingCycle(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingCycle) predicate.BillingCycle {
	return predicate.BillingCycle(sql.NotPredicates(p))
}
