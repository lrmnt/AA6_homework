// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrmnt/AA6_homework/analytics/ent/billingoperations"
	"github.com/lrmnt/AA6_homework/analytics/ent/predicate"
	"github.com/lrmnt/AA6_homework/analytics/ent/user"
)

// BillingOperationsQuery is the builder for querying BillingOperations entities.
type BillingOperationsQuery struct {
	config
	ctx        *QueryContext
	order      []billingoperations.OrderOption
	inters     []Interceptor
	predicates []predicate.BillingOperations
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingOperationsQuery builder.
func (boq *BillingOperationsQuery) Where(ps ...predicate.BillingOperations) *BillingOperationsQuery {
	boq.predicates = append(boq.predicates, ps...)
	return boq
}

// Limit the number of records to be returned by this query.
func (boq *BillingOperationsQuery) Limit(limit int) *BillingOperationsQuery {
	boq.ctx.Limit = &limit
	return boq
}

// Offset to start from.
func (boq *BillingOperationsQuery) Offset(offset int) *BillingOperationsQuery {
	boq.ctx.Offset = &offset
	return boq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (boq *BillingOperationsQuery) Unique(unique bool) *BillingOperationsQuery {
	boq.ctx.Unique = &unique
	return boq
}

// Order specifies how the records should be ordered.
func (boq *BillingOperationsQuery) Order(o ...billingoperations.OrderOption) *BillingOperationsQuery {
	boq.order = append(boq.order, o...)
	return boq
}

// QueryUser chains the current query on the "user" edge.
func (boq *BillingOperationsQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: boq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := boq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := boq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billingoperations.Table, billingoperations.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, billingoperations.UserTable, billingoperations.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(boq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillingOperations entity from the query.
// Returns a *NotFoundError when no BillingOperations was found.
func (boq *BillingOperationsQuery) First(ctx context.Context) (*BillingOperations, error) {
	nodes, err := boq.Limit(1).All(setContextOp(ctx, boq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billingoperations.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (boq *BillingOperationsQuery) FirstX(ctx context.Context) *BillingOperations {
	node, err := boq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingOperations ID from the query.
// Returns a *NotFoundError when no BillingOperations ID was found.
func (boq *BillingOperationsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = boq.Limit(1).IDs(setContextOp(ctx, boq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billingoperations.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (boq *BillingOperationsQuery) FirstIDX(ctx context.Context) int {
	id, err := boq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingOperations entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingOperations entity is found.
// Returns a *NotFoundError when no BillingOperations entities are found.
func (boq *BillingOperationsQuery) Only(ctx context.Context) (*BillingOperations, error) {
	nodes, err := boq.Limit(2).All(setContextOp(ctx, boq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billingoperations.Label}
	default:
		return nil, &NotSingularError{billingoperations.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (boq *BillingOperationsQuery) OnlyX(ctx context.Context) *BillingOperations {
	node, err := boq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingOperations ID in the query.
// Returns a *NotSingularError when more than one BillingOperations ID is found.
// Returns a *NotFoundError when no entities are found.
func (boq *BillingOperationsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = boq.Limit(2).IDs(setContextOp(ctx, boq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billingoperations.Label}
	default:
		err = &NotSingularError{billingoperations.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (boq *BillingOperationsQuery) OnlyIDX(ctx context.Context) int {
	id, err := boq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingOperationsSlice.
func (boq *BillingOperationsQuery) All(ctx context.Context) ([]*BillingOperations, error) {
	ctx = setContextOp(ctx, boq.ctx, "All")
	if err := boq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingOperations, *BillingOperationsQuery]()
	return withInterceptors[[]*BillingOperations](ctx, boq, qr, boq.inters)
}

// AllX is like All, but panics if an error occurs.
func (boq *BillingOperationsQuery) AllX(ctx context.Context) []*BillingOperations {
	nodes, err := boq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingOperations IDs.
func (boq *BillingOperationsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if boq.ctx.Unique == nil && boq.path != nil {
		boq.Unique(true)
	}
	ctx = setContextOp(ctx, boq.ctx, "IDs")
	if err = boq.Select(billingoperations.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (boq *BillingOperationsQuery) IDsX(ctx context.Context) []int {
	ids, err := boq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (boq *BillingOperationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, boq.ctx, "Count")
	if err := boq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, boq, querierCount[*BillingOperationsQuery](), boq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (boq *BillingOperationsQuery) CountX(ctx context.Context) int {
	count, err := boq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (boq *BillingOperationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, boq.ctx, "Exist")
	switch _, err := boq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (boq *BillingOperationsQuery) ExistX(ctx context.Context) bool {
	exist, err := boq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingOperationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (boq *BillingOperationsQuery) Clone() *BillingOperationsQuery {
	if boq == nil {
		return nil
	}
	return &BillingOperationsQuery{
		config:     boq.config,
		ctx:        boq.ctx.Clone(),
		order:      append([]billingoperations.OrderOption{}, boq.order...),
		inters:     append([]Interceptor{}, boq.inters...),
		predicates: append([]predicate.BillingOperations{}, boq.predicates...),
		withUser:   boq.withUser.Clone(),
		// clone intermediate query.
		sql:  boq.sql.Clone(),
		path: boq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (boq *BillingOperationsQuery) WithUser(opts ...func(*UserQuery)) *BillingOperationsQuery {
	query := (&UserClient{config: boq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	boq.withUser = query
	return boq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillingOperations.Query().
//		GroupBy(billingoperations.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (boq *BillingOperationsQuery) GroupBy(field string, fields ...string) *BillingOperationsGroupBy {
	boq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingOperationsGroupBy{build: boq}
	grbuild.flds = &boq.ctx.Fields
	grbuild.label = billingoperations.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.BillingOperations.Query().
//		Select(billingoperations.FieldUUID).
//		Scan(ctx, &v)
func (boq *BillingOperationsQuery) Select(fields ...string) *BillingOperationsSelect {
	boq.ctx.Fields = append(boq.ctx.Fields, fields...)
	sbuild := &BillingOperationsSelect{BillingOperationsQuery: boq}
	sbuild.label = billingoperations.Label
	sbuild.flds, sbuild.scan = &boq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingOperationsSelect configured with the given aggregations.
func (boq *BillingOperationsQuery) Aggregate(fns ...AggregateFunc) *BillingOperationsSelect {
	return boq.Select().Aggregate(fns...)
}

func (boq *BillingOperationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range boq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, boq); err != nil {
				return err
			}
		}
	}
	for _, f := range boq.ctx.Fields {
		if !billingoperations.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if boq.path != nil {
		prev, err := boq.path(ctx)
		if err != nil {
			return err
		}
		boq.sql = prev
	}
	return nil
}

func (boq *BillingOperationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingOperations, error) {
	var (
		nodes       = []*BillingOperations{}
		withFKs     = boq.withFKs
		_spec       = boq.querySpec()
		loadedTypes = [1]bool{
			boq.withUser != nil,
		}
	)
	if boq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, billingoperations.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingOperations).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingOperations{config: boq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, boq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := boq.withUser; query != nil {
		if err := boq.loadUser(ctx, query, nodes, nil,
			func(n *BillingOperations, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (boq *BillingOperationsQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*BillingOperations, init func(*BillingOperations), assign func(*BillingOperations, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BillingOperations)
	for i := range nodes {
		if nodes[i].billing_operations_user == nil {
			continue
		}
		fk := *nodes[i].billing_operations_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "billing_operations_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (boq *BillingOperationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := boq.querySpec()
	_spec.Node.Columns = boq.ctx.Fields
	if len(boq.ctx.Fields) > 0 {
		_spec.Unique = boq.ctx.Unique != nil && *boq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, boq.driver, _spec)
}

func (boq *BillingOperationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billingoperations.Table, billingoperations.Columns, sqlgraph.NewFieldSpec(billingoperations.FieldID, field.TypeInt))
	_spec.From = boq.sql
	if unique := boq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if boq.path != nil {
		_spec.Unique = true
	}
	if fields := boq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billingoperations.FieldID)
		for i := range fields {
			if fields[i] != billingoperations.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := boq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := boq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := boq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := boq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (boq *BillingOperationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(boq.driver.Dialect())
	t1 := builder.Table(billingoperations.Table)
	columns := boq.ctx.Fields
	if len(columns) == 0 {
		columns = billingoperations.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if boq.sql != nil {
		selector = boq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if boq.ctx.Unique != nil && *boq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range boq.predicates {
		p(selector)
	}
	for _, p := range boq.order {
		p(selector)
	}
	if offset := boq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := boq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BillingOperationsGroupBy is the group-by builder for BillingOperations entities.
type BillingOperationsGroupBy struct {
	selector
	build *BillingOperationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bogb *BillingOperationsGroupBy) Aggregate(fns ...AggregateFunc) *BillingOperationsGroupBy {
	bogb.fns = append(bogb.fns, fns...)
	return bogb
}

// Scan applies the selector query and scans the result into the given value.
func (bogb *BillingOperationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bogb.build.ctx, "GroupBy")
	if err := bogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingOperationsQuery, *BillingOperationsGroupBy](ctx, bogb.build, bogb, bogb.build.inters, v)
}

func (bogb *BillingOperationsGroupBy) sqlScan(ctx context.Context, root *BillingOperationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bogb.fns))
	for _, fn := range bogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bogb.flds)+len(bogb.fns))
		for _, f := range *bogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingOperationsSelect is the builder for selecting fields of BillingOperations entities.
type BillingOperationsSelect struct {
	*BillingOperationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bos *BillingOperationsSelect) Aggregate(fns ...AggregateFunc) *BillingOperationsSelect {
	bos.fns = append(bos.fns, fns...)
	return bos
}

// Scan applies the selector query and scans the result into the given value.
func (bos *BillingOperationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bos.ctx, "Select")
	if err := bos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingOperationsQuery, *BillingOperationsSelect](ctx, bos.BillingOperationsQuery, bos, bos.inters, v)
}

func (bos *BillingOperationsSelect) sqlScan(ctx context.Context, root *BillingOperationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bos.fns))
	for _, fn := range bos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
