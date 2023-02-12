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
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/upscaleoutput"
)

// UpscaleOutputUpdate is the builder for updating UpscaleOutput entities.
type UpscaleOutputUpdate struct {
	config
	hooks     []Hook
	mutation  *UpscaleOutputMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UpscaleOutputUpdate builder.
func (uou *UpscaleOutputUpdate) Where(ps ...predicate.UpscaleOutput) *UpscaleOutputUpdate {
	uou.mutation.Where(ps...)
	return uou
}

// SetImageURL sets the "image_url" field.
func (uou *UpscaleOutputUpdate) SetImageURL(s string) *UpscaleOutputUpdate {
	uou.mutation.SetImageURL(s)
	return uou
}

// SetUpscaleID sets the "upscale_id" field.
func (uou *UpscaleOutputUpdate) SetUpscaleID(u uuid.UUID) *UpscaleOutputUpdate {
	uou.mutation.SetUpscaleID(u)
	return uou
}

// SetUpdatedAt sets the "updated_at" field.
func (uou *UpscaleOutputUpdate) SetUpdatedAt(t time.Time) *UpscaleOutputUpdate {
	uou.mutation.SetUpdatedAt(t)
	return uou
}

// SetUpscalesID sets the "upscales" edge to the Upscale entity by ID.
func (uou *UpscaleOutputUpdate) SetUpscalesID(id uuid.UUID) *UpscaleOutputUpdate {
	uou.mutation.SetUpscalesID(id)
	return uou
}

// SetUpscales sets the "upscales" edge to the Upscale entity.
func (uou *UpscaleOutputUpdate) SetUpscales(u *Upscale) *UpscaleOutputUpdate {
	return uou.SetUpscalesID(u.ID)
}

// Mutation returns the UpscaleOutputMutation object of the builder.
func (uou *UpscaleOutputUpdate) Mutation() *UpscaleOutputMutation {
	return uou.mutation
}

// ClearUpscales clears the "upscales" edge to the Upscale entity.
func (uou *UpscaleOutputUpdate) ClearUpscales() *UpscaleOutputUpdate {
	uou.mutation.ClearUpscales()
	return uou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uou *UpscaleOutputUpdate) Save(ctx context.Context) (int, error) {
	uou.defaults()
	return withHooks[int, UpscaleOutputMutation](ctx, uou.sqlSave, uou.mutation, uou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uou *UpscaleOutputUpdate) SaveX(ctx context.Context) int {
	affected, err := uou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uou *UpscaleOutputUpdate) Exec(ctx context.Context) error {
	_, err := uou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uou *UpscaleOutputUpdate) ExecX(ctx context.Context) {
	if err := uou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uou *UpscaleOutputUpdate) defaults() {
	if _, ok := uou.mutation.UpdatedAt(); !ok {
		v := upscaleoutput.UpdateDefaultUpdatedAt()
		uou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uou *UpscaleOutputUpdate) check() error {
	if _, ok := uou.mutation.UpscalesID(); uou.mutation.UpscalesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UpscaleOutput.upscales"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uou *UpscaleOutputUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UpscaleOutputUpdate {
	uou.modifiers = append(uou.modifiers, modifiers...)
	return uou
}

func (uou *UpscaleOutputUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uou.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscaleoutput.Table,
			Columns: upscaleoutput.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscaleoutput.FieldID,
			},
		},
	}
	if ps := uou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uou.mutation.ImageURL(); ok {
		_spec.SetField(upscaleoutput.FieldImageURL, field.TypeString, value)
	}
	if value, ok := uou.mutation.UpdatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldUpdatedAt, field.TypeTime, value)
	}
	if uou.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uou.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uou.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscaleoutput.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uou.mutation.done = true
	return n, nil
}

// UpscaleOutputUpdateOne is the builder for updating a single UpscaleOutput entity.
type UpscaleOutputUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UpscaleOutputMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetImageURL sets the "image_url" field.
func (uouo *UpscaleOutputUpdateOne) SetImageURL(s string) *UpscaleOutputUpdateOne {
	uouo.mutation.SetImageURL(s)
	return uouo
}

// SetUpscaleID sets the "upscale_id" field.
func (uouo *UpscaleOutputUpdateOne) SetUpscaleID(u uuid.UUID) *UpscaleOutputUpdateOne {
	uouo.mutation.SetUpscaleID(u)
	return uouo
}

// SetUpdatedAt sets the "updated_at" field.
func (uouo *UpscaleOutputUpdateOne) SetUpdatedAt(t time.Time) *UpscaleOutputUpdateOne {
	uouo.mutation.SetUpdatedAt(t)
	return uouo
}

// SetUpscalesID sets the "upscales" edge to the Upscale entity by ID.
func (uouo *UpscaleOutputUpdateOne) SetUpscalesID(id uuid.UUID) *UpscaleOutputUpdateOne {
	uouo.mutation.SetUpscalesID(id)
	return uouo
}

// SetUpscales sets the "upscales" edge to the Upscale entity.
func (uouo *UpscaleOutputUpdateOne) SetUpscales(u *Upscale) *UpscaleOutputUpdateOne {
	return uouo.SetUpscalesID(u.ID)
}

// Mutation returns the UpscaleOutputMutation object of the builder.
func (uouo *UpscaleOutputUpdateOne) Mutation() *UpscaleOutputMutation {
	return uouo.mutation
}

// ClearUpscales clears the "upscales" edge to the Upscale entity.
func (uouo *UpscaleOutputUpdateOne) ClearUpscales() *UpscaleOutputUpdateOne {
	uouo.mutation.ClearUpscales()
	return uouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uouo *UpscaleOutputUpdateOne) Select(field string, fields ...string) *UpscaleOutputUpdateOne {
	uouo.fields = append([]string{field}, fields...)
	return uouo
}

// Save executes the query and returns the updated UpscaleOutput entity.
func (uouo *UpscaleOutputUpdateOne) Save(ctx context.Context) (*UpscaleOutput, error) {
	uouo.defaults()
	return withHooks[*UpscaleOutput, UpscaleOutputMutation](ctx, uouo.sqlSave, uouo.mutation, uouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uouo *UpscaleOutputUpdateOne) SaveX(ctx context.Context) *UpscaleOutput {
	node, err := uouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uouo *UpscaleOutputUpdateOne) Exec(ctx context.Context) error {
	_, err := uouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uouo *UpscaleOutputUpdateOne) ExecX(ctx context.Context) {
	if err := uouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uouo *UpscaleOutputUpdateOne) defaults() {
	if _, ok := uouo.mutation.UpdatedAt(); !ok {
		v := upscaleoutput.UpdateDefaultUpdatedAt()
		uouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uouo *UpscaleOutputUpdateOne) check() error {
	if _, ok := uouo.mutation.UpscalesID(); uouo.mutation.UpscalesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UpscaleOutput.upscales"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uouo *UpscaleOutputUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UpscaleOutputUpdateOne {
	uouo.modifiers = append(uouo.modifiers, modifiers...)
	return uouo
}

func (uouo *UpscaleOutputUpdateOne) sqlSave(ctx context.Context) (_node *UpscaleOutput, err error) {
	if err := uouo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscaleoutput.Table,
			Columns: upscaleoutput.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscaleoutput.FieldID,
			},
		},
	}
	id, ok := uouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UpscaleOutput.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscaleoutput.FieldID)
		for _, f := range fields {
			if !upscaleoutput.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != upscaleoutput.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uouo.mutation.ImageURL(); ok {
		_spec.SetField(upscaleoutput.FieldImageURL, field.TypeString, value)
	}
	if value, ok := uouo.mutation.UpdatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldUpdatedAt, field.TypeTime, value)
	}
	if uouo.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uouo.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uouo.modifiers...)
	_node = &UpscaleOutput{config: uouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscaleoutput.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uouo.mutation.done = true
	return _node, nil
}
