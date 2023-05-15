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
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/user"
)

// ApiTokenUpdate is the builder for updating ApiToken entities.
type ApiTokenUpdate struct {
	config
	hooks     []Hook
	mutation  *ApiTokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApiTokenUpdate builder.
func (atu *ApiTokenUpdate) Where(ps ...predicate.ApiToken) *ApiTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetHashedToken sets the "hashed_token" field.
func (atu *ApiTokenUpdate) SetHashedToken(s string) *ApiTokenUpdate {
	atu.mutation.SetHashedToken(s)
	return atu
}

// SetName sets the "name" field.
func (atu *ApiTokenUpdate) SetName(s string) *ApiTokenUpdate {
	atu.mutation.SetName(s)
	return atu
}

// SetShortString sets the "short_string" field.
func (atu *ApiTokenUpdate) SetShortString(s string) *ApiTokenUpdate {
	atu.mutation.SetShortString(s)
	return atu
}

// SetIsActive sets the "is_active" field.
func (atu *ApiTokenUpdate) SetIsActive(b bool) *ApiTokenUpdate {
	atu.mutation.SetIsActive(b)
	return atu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (atu *ApiTokenUpdate) SetNillableIsActive(b *bool) *ApiTokenUpdate {
	if b != nil {
		atu.SetIsActive(*b)
	}
	return atu
}

// SetUses sets the "uses" field.
func (atu *ApiTokenUpdate) SetUses(i int) *ApiTokenUpdate {
	atu.mutation.ResetUses()
	atu.mutation.SetUses(i)
	return atu
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (atu *ApiTokenUpdate) SetNillableUses(i *int) *ApiTokenUpdate {
	if i != nil {
		atu.SetUses(*i)
	}
	return atu
}

// AddUses adds i to the "uses" field.
func (atu *ApiTokenUpdate) AddUses(i int) *ApiTokenUpdate {
	atu.mutation.AddUses(i)
	return atu
}

// SetUserID sets the "user_id" field.
func (atu *ApiTokenUpdate) SetUserID(u uuid.UUID) *ApiTokenUpdate {
	atu.mutation.SetUserID(u)
	return atu
}

// SetLastUsedAt sets the "last_used_at" field.
func (atu *ApiTokenUpdate) SetLastUsedAt(t time.Time) *ApiTokenUpdate {
	atu.mutation.SetLastUsedAt(t)
	return atu
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (atu *ApiTokenUpdate) SetNillableLastUsedAt(t *time.Time) *ApiTokenUpdate {
	if t != nil {
		atu.SetLastUsedAt(*t)
	}
	return atu
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (atu *ApiTokenUpdate) ClearLastUsedAt() *ApiTokenUpdate {
	atu.mutation.ClearLastUsedAt()
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *ApiTokenUpdate) SetUpdatedAt(t time.Time) *ApiTokenUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *ApiTokenUpdate) SetUser(u *User) *ApiTokenUpdate {
	return atu.SetUserID(u.ID)
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (atu *ApiTokenUpdate) AddGenerationIDs(ids ...uuid.UUID) *ApiTokenUpdate {
	atu.mutation.AddGenerationIDs(ids...)
	return atu
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (atu *ApiTokenUpdate) AddGenerations(g ...*Generation) *ApiTokenUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return atu.AddGenerationIDs(ids...)
}

// AddUpscaleIDs adds the "upscales" edge to the Upscale entity by IDs.
func (atu *ApiTokenUpdate) AddUpscaleIDs(ids ...uuid.UUID) *ApiTokenUpdate {
	atu.mutation.AddUpscaleIDs(ids...)
	return atu
}

// AddUpscales adds the "upscales" edges to the Upscale entity.
func (atu *ApiTokenUpdate) AddUpscales(u ...*Upscale) *ApiTokenUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atu.AddUpscaleIDs(ids...)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atu *ApiTokenUpdate) Mutation() *ApiTokenMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *ApiTokenUpdate) ClearUser() *ApiTokenUpdate {
	atu.mutation.ClearUser()
	return atu
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (atu *ApiTokenUpdate) ClearGenerations() *ApiTokenUpdate {
	atu.mutation.ClearGenerations()
	return atu
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (atu *ApiTokenUpdate) RemoveGenerationIDs(ids ...uuid.UUID) *ApiTokenUpdate {
	atu.mutation.RemoveGenerationIDs(ids...)
	return atu
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (atu *ApiTokenUpdate) RemoveGenerations(g ...*Generation) *ApiTokenUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return atu.RemoveGenerationIDs(ids...)
}

// ClearUpscales clears all "upscales" edges to the Upscale entity.
func (atu *ApiTokenUpdate) ClearUpscales() *ApiTokenUpdate {
	atu.mutation.ClearUpscales()
	return atu
}

// RemoveUpscaleIDs removes the "upscales" edge to Upscale entities by IDs.
func (atu *ApiTokenUpdate) RemoveUpscaleIDs(ids ...uuid.UUID) *ApiTokenUpdate {
	atu.mutation.RemoveUpscaleIDs(ids...)
	return atu
}

// RemoveUpscales removes "upscales" edges to Upscale entities.
func (atu *ApiTokenUpdate) RemoveUpscales(u ...*Upscale) *ApiTokenUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atu.RemoveUpscaleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ApiTokenUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks[int, ApiTokenMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ApiTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ApiTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ApiTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *ApiTokenUpdate) defaults() {
	if _, ok := atu.mutation.UpdatedAt(); !ok {
		v := apitoken.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ApiTokenUpdate) check() error {
	if _, ok := atu.mutation.UserID(); atu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiToken.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (atu *ApiTokenUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApiTokenUpdate {
	atu.modifiers = append(atu.modifiers, modifiers...)
	return atu
}

func (atu *ApiTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apitoken.Table,
			Columns: apitoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apitoken.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.HashedToken(); ok {
		_spec.SetField(apitoken.FieldHashedToken, field.TypeString, value)
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
	}
	if value, ok := atu.mutation.ShortString(); ok {
		_spec.SetField(apitoken.FieldShortString, field.TypeString, value)
	}
	if value, ok := atu.mutation.IsActive(); ok {
		_spec.SetField(apitoken.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := atu.mutation.Uses(); ok {
		_spec.SetField(apitoken.FieldUses, field.TypeInt, value)
	}
	if value, ok := atu.mutation.AddedUses(); ok {
		_spec.AddField(apitoken.FieldUses, field.TypeInt, value)
	}
	if value, ok := atu.mutation.LastUsedAt(); ok {
		_spec.SetField(apitoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if atu.mutation.LastUsedAtCleared() {
		_spec.ClearField(apitoken.FieldLastUsedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if nodes := atu.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !atu.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if nodes := atu.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if atu.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
	if nodes := atu.mutation.RemovedUpscalesIDs(); len(nodes) > 0 && !atu.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
	_spec.AddModifiers(atu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// ApiTokenUpdateOne is the builder for updating a single ApiToken entity.
type ApiTokenUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApiTokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetHashedToken sets the "hashed_token" field.
func (atuo *ApiTokenUpdateOne) SetHashedToken(s string) *ApiTokenUpdateOne {
	atuo.mutation.SetHashedToken(s)
	return atuo
}

// SetName sets the "name" field.
func (atuo *ApiTokenUpdateOne) SetName(s string) *ApiTokenUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// SetShortString sets the "short_string" field.
func (atuo *ApiTokenUpdateOne) SetShortString(s string) *ApiTokenUpdateOne {
	atuo.mutation.SetShortString(s)
	return atuo
}

// SetIsActive sets the "is_active" field.
func (atuo *ApiTokenUpdateOne) SetIsActive(b bool) *ApiTokenUpdateOne {
	atuo.mutation.SetIsActive(b)
	return atuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (atuo *ApiTokenUpdateOne) SetNillableIsActive(b *bool) *ApiTokenUpdateOne {
	if b != nil {
		atuo.SetIsActive(*b)
	}
	return atuo
}

// SetUses sets the "uses" field.
func (atuo *ApiTokenUpdateOne) SetUses(i int) *ApiTokenUpdateOne {
	atuo.mutation.ResetUses()
	atuo.mutation.SetUses(i)
	return atuo
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (atuo *ApiTokenUpdateOne) SetNillableUses(i *int) *ApiTokenUpdateOne {
	if i != nil {
		atuo.SetUses(*i)
	}
	return atuo
}

// AddUses adds i to the "uses" field.
func (atuo *ApiTokenUpdateOne) AddUses(i int) *ApiTokenUpdateOne {
	atuo.mutation.AddUses(i)
	return atuo
}

// SetUserID sets the "user_id" field.
func (atuo *ApiTokenUpdateOne) SetUserID(u uuid.UUID) *ApiTokenUpdateOne {
	atuo.mutation.SetUserID(u)
	return atuo
}

// SetLastUsedAt sets the "last_used_at" field.
func (atuo *ApiTokenUpdateOne) SetLastUsedAt(t time.Time) *ApiTokenUpdateOne {
	atuo.mutation.SetLastUsedAt(t)
	return atuo
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (atuo *ApiTokenUpdateOne) SetNillableLastUsedAt(t *time.Time) *ApiTokenUpdateOne {
	if t != nil {
		atuo.SetLastUsedAt(*t)
	}
	return atuo
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (atuo *ApiTokenUpdateOne) ClearLastUsedAt() *ApiTokenUpdateOne {
	atuo.mutation.ClearLastUsedAt()
	return atuo
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *ApiTokenUpdateOne) SetUpdatedAt(t time.Time) *ApiTokenUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *ApiTokenUpdateOne) SetUser(u *User) *ApiTokenUpdateOne {
	return atuo.SetUserID(u.ID)
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (atuo *ApiTokenUpdateOne) AddGenerationIDs(ids ...uuid.UUID) *ApiTokenUpdateOne {
	atuo.mutation.AddGenerationIDs(ids...)
	return atuo
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (atuo *ApiTokenUpdateOne) AddGenerations(g ...*Generation) *ApiTokenUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return atuo.AddGenerationIDs(ids...)
}

// AddUpscaleIDs adds the "upscales" edge to the Upscale entity by IDs.
func (atuo *ApiTokenUpdateOne) AddUpscaleIDs(ids ...uuid.UUID) *ApiTokenUpdateOne {
	atuo.mutation.AddUpscaleIDs(ids...)
	return atuo
}

// AddUpscales adds the "upscales" edges to the Upscale entity.
func (atuo *ApiTokenUpdateOne) AddUpscales(u ...*Upscale) *ApiTokenUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atuo.AddUpscaleIDs(ids...)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atuo *ApiTokenUpdateOne) Mutation() *ApiTokenMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *ApiTokenUpdateOne) ClearUser() *ApiTokenUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// ClearGenerations clears all "generations" edges to the Generation entity.
func (atuo *ApiTokenUpdateOne) ClearGenerations() *ApiTokenUpdateOne {
	atuo.mutation.ClearGenerations()
	return atuo
}

// RemoveGenerationIDs removes the "generations" edge to Generation entities by IDs.
func (atuo *ApiTokenUpdateOne) RemoveGenerationIDs(ids ...uuid.UUID) *ApiTokenUpdateOne {
	atuo.mutation.RemoveGenerationIDs(ids...)
	return atuo
}

// RemoveGenerations removes "generations" edges to Generation entities.
func (atuo *ApiTokenUpdateOne) RemoveGenerations(g ...*Generation) *ApiTokenUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return atuo.RemoveGenerationIDs(ids...)
}

// ClearUpscales clears all "upscales" edges to the Upscale entity.
func (atuo *ApiTokenUpdateOne) ClearUpscales() *ApiTokenUpdateOne {
	atuo.mutation.ClearUpscales()
	return atuo
}

// RemoveUpscaleIDs removes the "upscales" edge to Upscale entities by IDs.
func (atuo *ApiTokenUpdateOne) RemoveUpscaleIDs(ids ...uuid.UUID) *ApiTokenUpdateOne {
	atuo.mutation.RemoveUpscaleIDs(ids...)
	return atuo
}

// RemoveUpscales removes "upscales" edges to Upscale entities.
func (atuo *ApiTokenUpdateOne) RemoveUpscales(u ...*Upscale) *ApiTokenUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atuo.RemoveUpscaleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ApiTokenUpdateOne) Select(field string, fields ...string) *ApiTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ApiToken entity.
func (atuo *ApiTokenUpdateOne) Save(ctx context.Context) (*ApiToken, error) {
	atuo.defaults()
	return withHooks[*ApiToken, ApiTokenMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) SaveX(ctx context.Context) *ApiToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ApiTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *ApiTokenUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdatedAt(); !ok {
		v := apitoken.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ApiTokenUpdateOne) check() error {
	if _, ok := atuo.mutation.UserID(); atuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiToken.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (atuo *ApiTokenUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApiTokenUpdateOne {
	atuo.modifiers = append(atuo.modifiers, modifiers...)
	return atuo
}

func (atuo *ApiTokenUpdateOne) sqlSave(ctx context.Context) (_node *ApiToken, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apitoken.Table,
			Columns: apitoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apitoken.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apitoken.FieldID)
		for _, f := range fields {
			if !apitoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.HashedToken(); ok {
		_spec.SetField(apitoken.FieldHashedToken, field.TypeString, value)
	}
	if value, ok := atuo.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
	}
	if value, ok := atuo.mutation.ShortString(); ok {
		_spec.SetField(apitoken.FieldShortString, field.TypeString, value)
	}
	if value, ok := atuo.mutation.IsActive(); ok {
		_spec.SetField(apitoken.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := atuo.mutation.Uses(); ok {
		_spec.SetField(apitoken.FieldUses, field.TypeInt, value)
	}
	if value, ok := atuo.mutation.AddedUses(); ok {
		_spec.AddField(apitoken.FieldUses, field.TypeInt, value)
	}
	if value, ok := atuo.mutation.LastUsedAt(); ok {
		_spec.SetField(apitoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if atuo.mutation.LastUsedAtCleared() {
		_spec.ClearField(apitoken.FieldLastUsedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if nodes := atuo.mutation.RemovedGenerationsIDs(); len(nodes) > 0 && !atuo.mutation.GenerationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if nodes := atuo.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
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
	if atuo.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
	if nodes := atuo.mutation.RemovedUpscalesIDs(); len(nodes) > 0 && !atuo.mutation.UpscalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
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
	_spec.AddModifiers(atuo.modifiers...)
	_node = &ApiToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
