// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/scheduler"
)

// SchedulerUpdate is the builder for updating Scheduler entities.
type SchedulerUpdate struct {
	config
	hooks     []Hook
	mutation  *SchedulerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SchedulerUpdate builder.
func (su *SchedulerUpdate) Where(ps ...predicate.Scheduler) *SchedulerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetNameInWorker sets the "name_in_worker" field.
func (su *SchedulerUpdate) SetNameInWorker(s string) *SchedulerUpdate {
	su.mutation.SetNameInWorker(s)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SchedulerUpdate) SetUpdatedAt(t time.Time) *SchedulerUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (su *SchedulerUpdate) AddGenerationIDs(ids ...uuid.UUID) *SchedulerUpdate {
	su.mutation.AddGenerationIDs(ids...)
	return su
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (su *SchedulerUpdate) AddGenerations(g ...*Generation) *SchedulerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.AddGenerationIDs(ids...)
}

// Mutation returns the SchedulerMutation object of the builder.
func (su *SchedulerUpdate) Mutation() *SchedulerMutation {
	return su.mutation
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (su *SchedulerUpdate) ClearGenerations() *SchedulerUpdate {
	su.mutation.ClearGenerations()
	return su
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (su *SchedulerUpdate) RemoveGenerationIDs(ids ...uuid.UUID) *SchedulerUpdate {
	su.mutation.RemoveGenerationIDs(ids...)
	return su
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (su *SchedulerUpdate) RemoveGenerations(g ...*Generation) *SchedulerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return su.RemoveGenerationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SchedulerUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks[int, SchedulerMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SchedulerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SchedulerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SchedulerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SchedulerUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := scheduler.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SchedulerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchedulerUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SchedulerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scheduler.Table,
			Columns: scheduler.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scheduler.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.NameInWorker(); ok {
		_spec.SetField(scheduler.FieldNameInWorker, field.TypeString, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(scheduler.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !su.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scheduler.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SchedulerUpdateOne is the builder for updating a single Scheduler entity.
type SchedulerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SchedulerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetNameInWorker sets the "name_in_worker" field.
func (suo *SchedulerUpdateOne) SetNameInWorker(s string) *SchedulerUpdateOne {
	suo.mutation.SetNameInWorker(s)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SchedulerUpdateOne) SetUpdatedAt(t time.Time) *SchedulerUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (suo *SchedulerUpdateOne) AddGenerationIDs(ids ...uuid.UUID) *SchedulerUpdateOne {
	suo.mutation.AddGenerationIDs(ids...)
	return suo
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (suo *SchedulerUpdateOne) AddGenerations(g ...*Generation) *SchedulerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.AddGenerationIDs(ids...)
}

// Mutation returns the SchedulerMutation object of the builder.
func (suo *SchedulerUpdateOne) Mutation() *SchedulerMutation {
	return suo.mutation
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (suo *SchedulerUpdateOne) ClearGenerations() *SchedulerUpdateOne {
	suo.mutation.ClearGenerations()
	return suo
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (suo *SchedulerUpdateOne) RemoveGenerationIDs(ids ...uuid.UUID) *SchedulerUpdateOne {
	suo.mutation.RemoveGenerationIDs(ids...)
	return suo
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (suo *SchedulerUpdateOne) RemoveGenerations(g ...*Generation) *SchedulerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return suo.RemoveGenerationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SchedulerUpdateOne) Select(field string, fields ...string) *SchedulerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Scheduler entity.
func (suo *SchedulerUpdateOne) Save(ctx context.Context) (*Scheduler, error) {
	suo.defaults()
	return withHooks[*Scheduler, SchedulerMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SchedulerUpdateOne) SaveX(ctx context.Context) *Scheduler {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SchedulerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SchedulerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SchedulerUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := scheduler.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SchedulerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchedulerUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SchedulerUpdateOne) sqlSave(ctx context.Context) (_node *Scheduler, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scheduler.Table,
			Columns: scheduler.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scheduler.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Scheduler.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scheduler.FieldID)
		for _, f := range fields {
			if !scheduler.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != scheduler.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.NameInWorker(); ok {
		_spec.SetField(scheduler.FieldNameInWorker, field.TypeString, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(scheduler.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !suo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scheduler.GenerationsTable,
			Columns: []string{scheduler.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Scheduler{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scheduler.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
