// Code generated by ent, DO NOT EDIT.

package operations

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Operations {
	return predicate.Operations(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Operations {
	return predicate.Operations(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Operations {
	return predicate.Operations(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Operations {
	return predicate.Operations(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Operations {
	return predicate.Operations(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Operations {
	return predicate.Operations(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Operations {
	return predicate.Operations(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldUUID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldAmount, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldTimestamp, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Operations {
	return predicate.Operations(sql.FieldLTE(FieldUUID, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.Operations {
	return predicate.Operations(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.Operations {
	return predicate.Operations(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.Operations {
	return predicate.Operations(sql.FieldLTE(FieldAmount, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Operations {
	return predicate.Operations(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Operations {
	return predicate.Operations(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Operations {
	return predicate.Operations(sql.FieldNotIn(FieldType, vs...))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Operations {
	return predicate.Operations(sql.FieldLTE(FieldTimestamp, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingCycle applies the HasEdge predicate on the "billing_cycle" edge.
func HasBillingCycle() predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BillingCycleTable, BillingCycleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingCycleWith applies the HasEdge predicate on the "billing_cycle" edge with a given conditions (other predicates).
func HasBillingCycleWith(preds ...predicate.BillingCycle) predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := newBillingCycleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.Operations {
	return predicate.Operations(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Operations) predicate.Operations {
	return predicate.Operations(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Operations) predicate.Operations {
	return predicate.Operations(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Operations) predicate.Operations {
	return predicate.Operations(sql.NotPredicates(p))
}
