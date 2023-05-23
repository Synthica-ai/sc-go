// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/deviceinfo"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/upscalemodel"
	"github.com/stablecog/sc-go/database/ent/upscaleoutput"
	"github.com/stablecog/sc-go/database/ent/user"
)

// UpscaleQuery is the builder for querying Upscale entities.
type UpscaleQuery struct {
	config
	ctx                *QueryContext
	order              []OrderFunc
	inters             []Interceptor
	predicates         []predicate.Upscale
	withUser           *UserQuery
	withDeviceInfo     *DeviceInfoQuery
	withUpscaleModels  *UpscaleModelQuery
	withAPITokens      *ApiTokenQuery
	withUpscaleOutputs *UpscaleOutputQuery
	modifiers          []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpscaleQuery builder.
func (uq *UpscaleQuery) Where(ps ...predicate.Upscale) *UpscaleQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UpscaleQuery) Limit(limit int) *UpscaleQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UpscaleQuery) Offset(offset int) *UpscaleQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UpscaleQuery) Unique(unique bool) *UpscaleQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UpscaleQuery) Order(o ...OrderFunc) *UpscaleQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryUser chains the current query on the "user" edge.
func (uq *UpscaleQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscale.Table, upscale.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upscale.UserTable, upscale.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDeviceInfo chains the current query on the "device_info" edge.
func (uq *UpscaleQuery) QueryDeviceInfo() *DeviceInfoQuery {
	query := (&DeviceInfoClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscale.Table, upscale.FieldID, selector),
			sqlgraph.To(deviceinfo.Table, deviceinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upscale.DeviceInfoTable, upscale.DeviceInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpscaleModels chains the current query on the "upscale_models" edge.
func (uq *UpscaleQuery) QueryUpscaleModels() *UpscaleModelQuery {
	query := (&UpscaleModelClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscale.Table, upscale.FieldID, selector),
			sqlgraph.To(upscalemodel.Table, upscalemodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upscale.UpscaleModelsTable, upscale.UpscaleModelsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAPITokens chains the current query on the "api_tokens" edge.
func (uq *UpscaleQuery) QueryAPITokens() *ApiTokenQuery {
	query := (&ApiTokenClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscale.Table, upscale.FieldID, selector),
			sqlgraph.To(apitoken.Table, apitoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upscale.APITokensTable, upscale.APITokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpscaleOutputs chains the current query on the "upscale_outputs" edge.
func (uq *UpscaleQuery) QueryUpscaleOutputs() *UpscaleOutputQuery {
	query := (&UpscaleOutputClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscale.Table, upscale.FieldID, selector),
			sqlgraph.To(upscaleoutput.Table, upscaleoutput.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, upscale.UpscaleOutputsTable, upscale.UpscaleOutputsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Upscale entity from the query.
// Returns a *NotFoundError when no Upscale was found.
func (uq *UpscaleQuery) First(ctx context.Context) (*Upscale, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upscale.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UpscaleQuery) FirstX(ctx context.Context) *Upscale {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Upscale ID from the query.
// Returns a *NotFoundError when no Upscale ID was found.
func (uq *UpscaleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upscale.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UpscaleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Upscale entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Upscale entity is found.
// Returns a *NotFoundError when no Upscale entities are found.
func (uq *UpscaleQuery) Only(ctx context.Context) (*Upscale, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upscale.Label}
	default:
		return nil, &NotSingularError{upscale.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UpscaleQuery) OnlyX(ctx context.Context) *Upscale {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Upscale ID in the query.
// Returns a *NotSingularError when more than one Upscale ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UpscaleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upscale.Label}
	default:
		err = &NotSingularError{upscale.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UpscaleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Upscales.
func (uq *UpscaleQuery) All(ctx context.Context) ([]*Upscale, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Upscale, *UpscaleQuery]()
	return withInterceptors[[]*Upscale](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UpscaleQuery) AllX(ctx context.Context) []*Upscale {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Upscale IDs.
func (uq *UpscaleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err := uq.Select(upscale.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UpscaleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UpscaleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UpscaleQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UpscaleQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UpscaleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UpscaleQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpscaleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UpscaleQuery) Clone() *UpscaleQuery {
	if uq == nil {
		return nil
	}
	return &UpscaleQuery{
		config:             uq.config,
		ctx:                uq.ctx.Clone(),
		order:              append([]OrderFunc{}, uq.order...),
		inters:             append([]Interceptor{}, uq.inters...),
		predicates:         append([]predicate.Upscale{}, uq.predicates...),
		withUser:           uq.withUser.Clone(),
		withDeviceInfo:     uq.withDeviceInfo.Clone(),
		withUpscaleModels:  uq.withUpscaleModels.Clone(),
		withAPITokens:      uq.withAPITokens.Clone(),
		withUpscaleOutputs: uq.withUpscaleOutputs.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpscaleQuery) WithUser(opts ...func(*UserQuery)) *UpscaleQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withUser = query
	return uq
}

// WithDeviceInfo tells the query-builder to eager-load the nodes that are connected to
// the "device_info" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpscaleQuery) WithDeviceInfo(opts ...func(*DeviceInfoQuery)) *UpscaleQuery {
	query := (&DeviceInfoClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withDeviceInfo = query
	return uq
}

// WithUpscaleModels tells the query-builder to eager-load the nodes that are connected to
// the "upscale_models" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpscaleQuery) WithUpscaleModels(opts ...func(*UpscaleModelQuery)) *UpscaleQuery {
	query := (&UpscaleModelClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withUpscaleModels = query
	return uq
}

// WithAPITokens tells the query-builder to eager-load the nodes that are connected to
// the "api_tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpscaleQuery) WithAPITokens(opts ...func(*ApiTokenQuery)) *UpscaleQuery {
	query := (&ApiTokenClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withAPITokens = query
	return uq
}

// WithUpscaleOutputs tells the query-builder to eager-load the nodes that are connected to
// the "upscale_outputs" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UpscaleQuery) WithUpscaleOutputs(opts ...func(*UpscaleOutputQuery)) *UpscaleQuery {
	query := (&UpscaleOutputClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withUpscaleOutputs = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Width int32 `json:"width,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Upscale.Query().
//		GroupBy(upscale.FieldWidth).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UpscaleQuery) GroupBy(field string, fields ...string) *UpscaleGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpscaleGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = upscale.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Width int32 `json:"width,omitempty"`
//	}
//
//	client.Upscale.Query().
//		Select(upscale.FieldWidth).
//		Scan(ctx, &v)
func (uq *UpscaleQuery) Select(fields ...string) *UpscaleSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UpscaleSelect{UpscaleQuery: uq}
	sbuild.label = upscale.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpscaleSelect configured with the given aggregations.
func (uq *UpscaleQuery) Aggregate(fns ...AggregateFunc) *UpscaleSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UpscaleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !upscale.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UpscaleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Upscale, error) {
	var (
		nodes       = []*Upscale{}
		_spec       = uq.querySpec()
		loadedTypes = [5]bool{
			uq.withUser != nil,
			uq.withDeviceInfo != nil,
			uq.withUpscaleModels != nil,
			uq.withAPITokens != nil,
			uq.withUpscaleOutputs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Upscale).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Upscale{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withUser; query != nil {
		if err := uq.loadUser(ctx, query, nodes, nil,
			func(n *Upscale, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withDeviceInfo; query != nil {
		if err := uq.loadDeviceInfo(ctx, query, nodes, nil,
			func(n *Upscale, e *DeviceInfo) { n.Edges.DeviceInfo = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withUpscaleModels; query != nil {
		if err := uq.loadUpscaleModels(ctx, query, nodes, nil,
			func(n *Upscale, e *UpscaleModel) { n.Edges.UpscaleModels = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withAPITokens; query != nil {
		if err := uq.loadAPITokens(ctx, query, nodes, nil,
			func(n *Upscale, e *ApiToken) { n.Edges.APITokens = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withUpscaleOutputs; query != nil {
		if err := uq.loadUpscaleOutputs(ctx, query, nodes,
			func(n *Upscale) { n.Edges.UpscaleOutputs = []*UpscaleOutput{} },
			func(n *Upscale, e *UpscaleOutput) { n.Edges.UpscaleOutputs = append(n.Edges.UpscaleOutputs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UpscaleQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Upscale, init func(*Upscale), assign func(*Upscale, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Upscale)
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
func (uq *UpscaleQuery) loadDeviceInfo(ctx context.Context, query *DeviceInfoQuery, nodes []*Upscale, init func(*Upscale), assign func(*Upscale, *DeviceInfo)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Upscale)
	for i := range nodes {
		fk := nodes[i].DeviceInfoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(deviceinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_info_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UpscaleQuery) loadUpscaleModels(ctx context.Context, query *UpscaleModelQuery, nodes []*Upscale, init func(*Upscale), assign func(*Upscale, *UpscaleModel)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Upscale)
	for i := range nodes {
		fk := nodes[i].ModelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(upscalemodel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UpscaleQuery) loadAPITokens(ctx context.Context, query *ApiTokenQuery, nodes []*Upscale, init func(*Upscale), assign func(*Upscale, *ApiToken)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Upscale)
	for i := range nodes {
		if nodes[i].APITokenID == nil {
			continue
		}
		fk := *nodes[i].APITokenID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(apitoken.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "api_token_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UpscaleQuery) loadUpscaleOutputs(ctx context.Context, query *UpscaleOutputQuery, nodes []*Upscale, init func(*Upscale), assign func(*Upscale, *UpscaleOutput)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Upscale)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.UpscaleOutput(func(s *sql.Selector) {
		s.Where(sql.InValues(upscale.UpscaleOutputsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UpscaleID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upscale_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UpscaleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UpscaleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscale.Table,
			Columns: upscale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscale.FieldID)
		for i := range fields {
			if fields[i] != upscale.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UpscaleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(upscale.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = upscale.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range uq.modifiers {
		m(selector)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uq *UpscaleQuery) Modify(modifiers ...func(s *sql.Selector)) *UpscaleSelect {
	uq.modifiers = append(uq.modifiers, modifiers...)
	return uq.Select()
}

// UpscaleGroupBy is the group-by builder for Upscale entities.
type UpscaleGroupBy struct {
	selector
	build *UpscaleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UpscaleGroupBy) Aggregate(fns ...AggregateFunc) *UpscaleGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UpscaleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleQuery, *UpscaleGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UpscaleGroupBy) sqlScan(ctx context.Context, root *UpscaleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpscaleSelect is the builder for selecting fields of Upscale entities.
type UpscaleSelect struct {
	*UpscaleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UpscaleSelect) Aggregate(fns ...AggregateFunc) *UpscaleSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UpscaleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleQuery, *UpscaleSelect](ctx, us.UpscaleQuery, us, us.inters, v)
}

func (us *UpscaleSelect) sqlScan(ctx context.Context, root *UpscaleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (us *UpscaleSelect) Modify(modifiers ...func(s *sql.Selector)) *UpscaleSelect {
	us.modifiers = append(us.modifiers, modifiers...)
	return us
}
