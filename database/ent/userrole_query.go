// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/database/ent/userrole"
)

// UserRoleQuery is the builder for querying UserRole entities.
type UserRoleQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.UserRole
	withUsers  *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserRoleQuery builder.
func (urq *UserRoleQuery) Where(ps ...predicate.UserRole) *UserRoleQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit the number of records to be returned by this query.
func (urq *UserRoleQuery) Limit(limit int) *UserRoleQuery {
	urq.ctx.Limit = &limit
	return urq
}

// Offset to start from.
func (urq *UserRoleQuery) Offset(offset int) *UserRoleQuery {
	urq.ctx.Offset = &offset
	return urq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urq *UserRoleQuery) Unique(unique bool) *UserRoleQuery {
	urq.ctx.Unique = &unique
	return urq
}

// Order specifies how the records should be ordered.
func (urq *UserRoleQuery) Order(o ...OrderFunc) *UserRoleQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// QueryUsers chains the current query on the "users" edge.
func (urq *UserRoleQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: urq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userrole.UsersTable, userrole.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserRole entity from the query.
// Returns a *NotFoundError when no UserRole was found.
func (urq *UserRoleQuery) First(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(1).All(setContextOp(ctx, urq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UserRoleQuery) FirstX(ctx context.Context) *UserRole {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserRole ID from the query.
// Returns a *NotFoundError when no UserRole ID was found.
func (urq *UserRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(1).IDs(setContextOp(ctx, urq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UserRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserRole entity is found.
// Returns a *NotFoundError when no UserRole entities are found.
func (urq *UserRoleQuery) Only(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(2).All(setContextOp(ctx, urq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userrole.Label}
	default:
		return nil, &NotSingularError{userrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyX(ctx context.Context) *UserRole {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserRole ID in the query.
// Returns a *NotSingularError when more than one UserRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (urq *UserRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(2).IDs(setContextOp(ctx, urq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = &NotSingularError{userrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserRoles.
func (urq *UserRoleQuery) All(ctx context.Context) ([]*UserRole, error) {
	ctx = setContextOp(ctx, urq.ctx, "All")
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserRole, *UserRoleQuery]()
	return withInterceptors[[]*UserRole](ctx, urq, qr, urq.inters)
}

// AllX is like All, but panics if an error occurs.
func (urq *UserRoleQuery) AllX(ctx context.Context) []*UserRole {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserRole IDs.
func (urq *UserRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, urq.ctx, "IDs")
	if err := urq.Select(userrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UserRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UserRoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, urq.ctx, "Count")
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, urq, querierCount[*UserRoleQuery](), urq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UserRoleQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UserRoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, urq.ctx, "Exist")
	switch _, err := urq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UserRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UserRoleQuery) Clone() *UserRoleQuery {
	if urq == nil {
		return nil
	}
	return &UserRoleQuery{
		config:     urq.config,
		ctx:        urq.ctx.Clone(),
		order:      append([]OrderFunc{}, urq.order...),
		inters:     append([]Interceptor{}, urq.inters...),
		predicates: append([]predicate.UserRole{}, urq.predicates...),
		withUsers:  urq.withUsers.Clone(),
		// clone intermediate query.
		sql:  urq.sql.Clone(),
		path: urq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (urq *UserRoleQuery) WithUsers(opts ...func(*UserQuery)) *UserRoleQuery {
	query := (&UserClient{config: urq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	urq.withUsers = query
	return urq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserRole.Query().
//		GroupBy(userrole.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (urq *UserRoleQuery) GroupBy(field string, fields ...string) *UserRoleGroupBy {
	urq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserRoleGroupBy{build: urq}
	grbuild.flds = &urq.ctx.Fields
	grbuild.label = userrole.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.UserRole.Query().
//		Select(userrole.FieldUserID).
//		Scan(ctx, &v)
func (urq *UserRoleQuery) Select(fields ...string) *UserRoleSelect {
	urq.ctx.Fields = append(urq.ctx.Fields, fields...)
	sbuild := &UserRoleSelect{UserRoleQuery: urq}
	sbuild.label = userrole.Label
	sbuild.flds, sbuild.scan = &urq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserRoleSelect configured with the given aggregations.
func (urq *UserRoleQuery) Aggregate(fns ...AggregateFunc) *UserRoleSelect {
	return urq.Select().Aggregate(fns...)
}

func (urq *UserRoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range urq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, urq); err != nil {
				return err
			}
		}
	}
	for _, f := range urq.ctx.Fields {
		if !userrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	return nil
}

func (urq *UserRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserRole, error) {
	var (
		nodes       = []*UserRole{}
		_spec       = urq.querySpec()
		loadedTypes = [1]bool{
			urq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserRole{config: urq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := urq.withUsers; query != nil {
		if err := urq.loadUsers(ctx, query, nodes, nil,
			func(n *UserRole, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (urq *UserRoleQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*UserRole, init func(*UserRole), assign func(*UserRole, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserRole)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (urq *UserRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	_spec.Node.Columns = urq.ctx.Fields
	if len(urq.ctx.Fields) > 0 {
		_spec.Unique = urq.ctx.Unique != nil && *urq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UserRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
		From:   urq.sql,
		Unique: true,
	}
	if unique := urq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := urq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrole.FieldID)
		for i := range fields {
			if fields[i] != userrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urq *UserRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(userrole.Table)
	columns := urq.ctx.Fields
	if len(columns) == 0 {
		columns = userrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urq.ctx.Unique != nil && *urq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range urq.modifiers {
		m(selector)
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector)
	}
	if offset := urq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (urq *UserRoleQuery) Modify(modifiers ...func(s *sql.Selector)) *UserRoleSelect {
	urq.modifiers = append(urq.modifiers, modifiers...)
	return urq.Select()
}

// UserRoleGroupBy is the group-by builder for UserRole entities.
type UserRoleGroupBy struct {
	selector
	build *UserRoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UserRoleGroupBy) Aggregate(fns ...AggregateFunc) *UserRoleGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the selector query and scans the result into the given value.
func (urgb *UserRoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urgb.build.ctx, "GroupBy")
	if err := urgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserRoleQuery, *UserRoleGroupBy](ctx, urgb.build, urgb, urgb.build.inters, v)
}

func (urgb *UserRoleGroupBy) sqlScan(ctx context.Context, root *UserRoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(urgb.fns))
	for _, fn := range urgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*urgb.flds)+len(urgb.fns))
		for _, f := range *urgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*urgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserRoleSelect is the builder for selecting fields of UserRole entities.
type UserRoleSelect struct {
	*UserRoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (urs *UserRoleSelect) Aggregate(fns ...AggregateFunc) *UserRoleSelect {
	urs.fns = append(urs.fns, fns...)
	return urs
}

// Scan applies the selector query and scans the result into the given value.
func (urs *UserRoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urs.ctx, "Select")
	if err := urs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserRoleQuery, *UserRoleSelect](ctx, urs.UserRoleQuery, urs, urs.inters, v)
}

func (urs *UserRoleSelect) sqlScan(ctx context.Context, root *UserRoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(urs.fns))
	for _, fn := range urs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*urs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (urs *UserRoleSelect) Modify(modifiers ...func(s *sql.Selector)) *UserRoleSelect {
	urs.modifiers = append(urs.modifiers, modifiers...)
	return urs
}
