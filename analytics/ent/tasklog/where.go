// Code generated by ent, DO NOT EDIT.

package tasklog

import (
	"entgo.io/ent/dialect/sql"
	"github.com/lrmnt/AA6_homework/analytics/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldID, id))
}

// IdempotencyKey applies equality check predicate on the "idempotency_key" field. It's identical to IdempotencyKeyEQ.
func IdempotencyKey(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldIdempotencyKey, v))
}

// IdempotencyKeyEQ applies the EQ predicate on the "idempotency_key" field.
func IdempotencyKeyEQ(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEQ(FieldIdempotencyKey, v))
}

// IdempotencyKeyNEQ applies the NEQ predicate on the "idempotency_key" field.
func IdempotencyKeyNEQ(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNEQ(FieldIdempotencyKey, v))
}

// IdempotencyKeyIn applies the In predicate on the "idempotency_key" field.
func IdempotencyKeyIn(vs ...string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldIn(FieldIdempotencyKey, vs...))
}

// IdempotencyKeyNotIn applies the NotIn predicate on the "idempotency_key" field.
func IdempotencyKeyNotIn(vs ...string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldNotIn(FieldIdempotencyKey, vs...))
}

// IdempotencyKeyGT applies the GT predicate on the "idempotency_key" field.
func IdempotencyKeyGT(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGT(FieldIdempotencyKey, v))
}

// IdempotencyKeyGTE applies the GTE predicate on the "idempotency_key" field.
func IdempotencyKeyGTE(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldGTE(FieldIdempotencyKey, v))
}

// IdempotencyKeyLT applies the LT predicate on the "idempotency_key" field.
func IdempotencyKeyLT(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLT(FieldIdempotencyKey, v))
}

// IdempotencyKeyLTE applies the LTE predicate on the "idempotency_key" field.
func IdempotencyKeyLTE(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldLTE(FieldIdempotencyKey, v))
}

// IdempotencyKeyContains applies the Contains predicate on the "idempotency_key" field.
func IdempotencyKeyContains(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldContains(FieldIdempotencyKey, v))
}

// IdempotencyKeyHasPrefix applies the HasPrefix predicate on the "idempotency_key" field.
func IdempotencyKeyHasPrefix(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldHasPrefix(FieldIdempotencyKey, v))
}

// IdempotencyKeyHasSuffix applies the HasSuffix predicate on the "idempotency_key" field.
func IdempotencyKeyHasSuffix(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldHasSuffix(FieldIdempotencyKey, v))
}

// IdempotencyKeyEqualFold applies the EqualFold predicate on the "idempotency_key" field.
func IdempotencyKeyEqualFold(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldEqualFold(FieldIdempotencyKey, v))
}

// IdempotencyKeyContainsFold applies the ContainsFold predicate on the "idempotency_key" field.
func IdempotencyKeyContainsFold(v string) predicate.TaskLog {
	return predicate.TaskLog(sql.FieldContainsFold(FieldIdempotencyKey, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskLog) predicate.TaskLog {
	return predicate.TaskLog(sql.NotPredicates(p))
}