// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/deviceinfo"
	"github.com/stablecog/go-apps/database/ent/generation"
	"github.com/stablecog/go-apps/database/ent/generationmodel"
	"github.com/stablecog/go-apps/database/ent/generationoutput"
	"github.com/stablecog/go-apps/database/ent/negativeprompt"
	"github.com/stablecog/go-apps/database/ent/prompt"
	"github.com/stablecog/go-apps/database/ent/scheduler"
	"github.com/stablecog/go-apps/database/ent/user"
)

// GenerationCreate is the builder for creating a Generation entity.
type GenerationCreate struct {
	config
	mutation *GenerationMutation
	hooks    []Hook
}

// SetWidth sets the "width" field.
func (gc *GenerationCreate) SetWidth(i int32) *GenerationCreate {
	gc.mutation.SetWidth(i)
	return gc
}

// SetHeight sets the "height" field.
func (gc *GenerationCreate) SetHeight(i int32) *GenerationCreate {
	gc.mutation.SetHeight(i)
	return gc
}

// SetNumInterferenceSteps sets the "num_interference_steps" field.
func (gc *GenerationCreate) SetNumInterferenceSteps(i int32) *GenerationCreate {
	gc.mutation.SetNumInterferenceSteps(i)
	return gc
}

// SetGuidanceScale sets the "guidance_scale" field.
func (gc *GenerationCreate) SetGuidanceScale(f float32) *GenerationCreate {
	gc.mutation.SetGuidanceScale(f)
	return gc
}

// SetSeed sets the "seed" field.
func (gc *GenerationCreate) SetSeed(i int) *GenerationCreate {
	gc.mutation.SetSeed(i)
	return gc
}

// SetStatus sets the "status" field.
func (gc *GenerationCreate) SetStatus(ge generation.Status) *GenerationCreate {
	gc.mutation.SetStatus(ge)
	return gc
}

// SetFailureReason sets the "failure_reason" field.
func (gc *GenerationCreate) SetFailureReason(s string) *GenerationCreate {
	gc.mutation.SetFailureReason(s)
	return gc
}

// SetNillableFailureReason sets the "failure_reason" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableFailureReason(s *string) *GenerationCreate {
	if s != nil {
		gc.SetFailureReason(*s)
	}
	return gc
}

// SetCountryCode sets the "country_code" field.
func (gc *GenerationCreate) SetCountryCode(s string) *GenerationCreate {
	gc.mutation.SetCountryCode(s)
	return gc
}

// SetGalleryStatus sets the "gallery_status" field.
func (gc *GenerationCreate) SetGalleryStatus(gs generation.GalleryStatus) *GenerationCreate {
	gc.mutation.SetGalleryStatus(gs)
	return gc
}

// SetNillableGalleryStatus sets the "gallery_status" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableGalleryStatus(gs *generation.GalleryStatus) *GenerationCreate {
	if gs != nil {
		gc.SetGalleryStatus(*gs)
	}
	return gc
}

// SetInitImageURL sets the "init_image_url" field.
func (gc *GenerationCreate) SetInitImageURL(s string) *GenerationCreate {
	gc.mutation.SetInitImageURL(s)
	return gc
}

// SetNillableInitImageURL sets the "init_image_url" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableInitImageURL(s *string) *GenerationCreate {
	if s != nil {
		gc.SetInitImageURL(*s)
	}
	return gc
}

// SetPromptID sets the "prompt_id" field.
func (gc *GenerationCreate) SetPromptID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetPromptID(u)
	return gc
}

// SetNegativePromptID sets the "negative_prompt_id" field.
func (gc *GenerationCreate) SetNegativePromptID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetNegativePromptID(u)
	return gc
}

// SetNillableNegativePromptID sets the "negative_prompt_id" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableNegativePromptID(u *uuid.UUID) *GenerationCreate {
	if u != nil {
		gc.SetNegativePromptID(*u)
	}
	return gc
}

// SetModelID sets the "model_id" field.
func (gc *GenerationCreate) SetModelID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetModelID(u)
	return gc
}

// SetSchedulerID sets the "scheduler_id" field.
func (gc *GenerationCreate) SetSchedulerID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetSchedulerID(u)
	return gc
}

// SetUserID sets the "user_id" field.
func (gc *GenerationCreate) SetUserID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetUserID(u)
	return gc
}

// SetDeviceInfoID sets the "device_info_id" field.
func (gc *GenerationCreate) SetDeviceInfoID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetDeviceInfoID(u)
	return gc
}

// SetStartedAt sets the "started_at" field.
func (gc *GenerationCreate) SetStartedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetStartedAt(t)
	return gc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableStartedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetStartedAt(*t)
	}
	return gc
}

// SetCompletedAt sets the "completed_at" field.
func (gc *GenerationCreate) SetCompletedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetCompletedAt(t)
	return gc
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableCompletedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetCompletedAt(*t)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GenerationCreate) SetCreatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableCreatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GenerationCreate) SetUpdatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableUpdatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GenerationCreate) SetID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetID(u)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableID(u *uuid.UUID) *GenerationCreate {
	if u != nil {
		gc.SetID(*u)
	}
	return gc
}

// SetDeviceInfo sets the "device_info" edge to the DeviceInfo entity.
func (gc *GenerationCreate) SetDeviceInfo(d *DeviceInfo) *GenerationCreate {
	return gc.SetDeviceInfoID(d.ID)
}

// SetScheduler sets the "scheduler" edge to the Scheduler entity.
func (gc *GenerationCreate) SetScheduler(s *Scheduler) *GenerationCreate {
	return gc.SetSchedulerID(s.ID)
}

// SetPrompt sets the "prompt" edge to the Prompt entity.
func (gc *GenerationCreate) SetPrompt(p *Prompt) *GenerationCreate {
	return gc.SetPromptID(p.ID)
}

// SetNegativePrompt sets the "negative_prompt" edge to the NegativePrompt entity.
func (gc *GenerationCreate) SetNegativePrompt(n *NegativePrompt) *GenerationCreate {
	return gc.SetNegativePromptID(n.ID)
}

// SetGenerationModelID sets the "generation_model" edge to the GenerationModel entity by ID.
func (gc *GenerationCreate) SetGenerationModelID(id uuid.UUID) *GenerationCreate {
	gc.mutation.SetGenerationModelID(id)
	return gc
}

// SetGenerationModel sets the "generation_model" edge to the GenerationModel entity.
func (gc *GenerationCreate) SetGenerationModel(g *GenerationModel) *GenerationCreate {
	return gc.SetGenerationModelID(g.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (gc *GenerationCreate) SetUsersID(id uuid.UUID) *GenerationCreate {
	gc.mutation.SetUsersID(id)
	return gc
}

// SetUsers sets the "users" edge to the User entity.
func (gc *GenerationCreate) SetUsers(u *User) *GenerationCreate {
	return gc.SetUsersID(u.ID)
}

// AddGenerationOutputIDs adds the "generation_outputs" edge to the GenerationOutput entity by IDs.
func (gc *GenerationCreate) AddGenerationOutputIDs(ids ...uuid.UUID) *GenerationCreate {
	gc.mutation.AddGenerationOutputIDs(ids...)
	return gc
}

// AddGenerationOutputs adds the "generation_outputs" edges to the GenerationOutput entity.
func (gc *GenerationCreate) AddGenerationOutputs(g ...*GenerationOutput) *GenerationCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddGenerationOutputIDs(ids...)
}

// Mutation returns the GenerationMutation object of the builder.
func (gc *GenerationCreate) Mutation() *GenerationMutation {
	return gc.mutation
}

// Save creates the Generation in the database.
func (gc *GenerationCreate) Save(ctx context.Context) (*Generation, error) {
	gc.defaults()
	return withHooks[*Generation, GenerationMutation](ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenerationCreate) SaveX(ctx context.Context) *Generation {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GenerationCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GenerationCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GenerationCreate) defaults() {
	if _, ok := gc.mutation.GalleryStatus(); !ok {
		v := generation.DefaultGalleryStatus
		gc.mutation.SetGalleryStatus(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := generation.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := generation.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := generation.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenerationCreate) check() error {
	if _, ok := gc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Generation.width"`)}
	}
	if _, ok := gc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Generation.height"`)}
	}
	if _, ok := gc.mutation.NumInterferenceSteps(); !ok {
		return &ValidationError{Name: "num_interference_steps", err: errors.New(`ent: missing required field "Generation.num_interference_steps"`)}
	}
	if _, ok := gc.mutation.GuidanceScale(); !ok {
		return &ValidationError{Name: "guidance_scale", err: errors.New(`ent: missing required field "Generation.guidance_scale"`)}
	}
	if _, ok := gc.mutation.Seed(); !ok {
		return &ValidationError{Name: "seed", err: errors.New(`ent: missing required field "Generation.seed"`)}
	}
	if _, ok := gc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Generation.status"`)}
	}
	if v, ok := gc.mutation.Status(); ok {
		if err := generation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Generation.status": %w`, err)}
		}
	}
	if _, ok := gc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`ent: missing required field "Generation.country_code"`)}
	}
	if _, ok := gc.mutation.GalleryStatus(); !ok {
		return &ValidationError{Name: "gallery_status", err: errors.New(`ent: missing required field "Generation.gallery_status"`)}
	}
	if v, ok := gc.mutation.GalleryStatus(); ok {
		if err := generation.GalleryStatusValidator(v); err != nil {
			return &ValidationError{Name: "gallery_status", err: fmt.Errorf(`ent: validator failed for field "Generation.gallery_status": %w`, err)}
		}
	}
	if _, ok := gc.mutation.PromptID(); !ok {
		return &ValidationError{Name: "prompt_id", err: errors.New(`ent: missing required field "Generation.prompt_id"`)}
	}
	if _, ok := gc.mutation.ModelID(); !ok {
		return &ValidationError{Name: "model_id", err: errors.New(`ent: missing required field "Generation.model_id"`)}
	}
	if _, ok := gc.mutation.SchedulerID(); !ok {
		return &ValidationError{Name: "scheduler_id", err: errors.New(`ent: missing required field "Generation.scheduler_id"`)}
	}
	if _, ok := gc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Generation.user_id"`)}
	}
	if _, ok := gc.mutation.DeviceInfoID(); !ok {
		return &ValidationError{Name: "device_info_id", err: errors.New(`ent: missing required field "Generation.device_info_id"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Generation.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Generation.updated_at"`)}
	}
	if _, ok := gc.mutation.DeviceInfoID(); !ok {
		return &ValidationError{Name: "device_info", err: errors.New(`ent: missing required edge "Generation.device_info"`)}
	}
	if _, ok := gc.mutation.SchedulerID(); !ok {
		return &ValidationError{Name: "scheduler", err: errors.New(`ent: missing required edge "Generation.scheduler"`)}
	}
	if _, ok := gc.mutation.PromptID(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required edge "Generation.prompt"`)}
	}
	if _, ok := gc.mutation.GenerationModelID(); !ok {
		return &ValidationError{Name: "generation_model", err: errors.New(`ent: missing required edge "Generation.generation_model"`)}
	}
	if _, ok := gc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Generation.users"`)}
	}
	return nil
}

func (gc *GenerationCreate) sqlSave(ctx context.Context) (*Generation, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GenerationCreate) createSpec() (*Generation, *sqlgraph.CreateSpec) {
	var (
		_node = &Generation{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: generation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generation.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gc.mutation.Width(); ok {
		_spec.SetField(generation.FieldWidth, field.TypeInt32, value)
		_node.Width = value
	}
	if value, ok := gc.mutation.Height(); ok {
		_spec.SetField(generation.FieldHeight, field.TypeInt32, value)
		_node.Height = value
	}
	if value, ok := gc.mutation.NumInterferenceSteps(); ok {
		_spec.SetField(generation.FieldNumInterferenceSteps, field.TypeInt32, value)
		_node.NumInterferenceSteps = value
	}
	if value, ok := gc.mutation.GuidanceScale(); ok {
		_spec.SetField(generation.FieldGuidanceScale, field.TypeFloat32, value)
		_node.GuidanceScale = value
	}
	if value, ok := gc.mutation.Seed(); ok {
		_spec.SetField(generation.FieldSeed, field.TypeInt, value)
		_node.Seed = value
	}
	if value, ok := gc.mutation.Status(); ok {
		_spec.SetField(generation.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := gc.mutation.FailureReason(); ok {
		_spec.SetField(generation.FieldFailureReason, field.TypeString, value)
		_node.FailureReason = &value
	}
	if value, ok := gc.mutation.CountryCode(); ok {
		_spec.SetField(generation.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if value, ok := gc.mutation.GalleryStatus(); ok {
		_spec.SetField(generation.FieldGalleryStatus, field.TypeEnum, value)
		_node.GalleryStatus = value
	}
	if value, ok := gc.mutation.InitImageURL(); ok {
		_spec.SetField(generation.FieldInitImageURL, field.TypeString, value)
		_node.InitImageURL = &value
	}
	if value, ok := gc.mutation.StartedAt(); ok {
		_spec.SetField(generation.FieldStartedAt, field.TypeTime, value)
		_node.StartedAt = &value
	}
	if value, ok := gc.mutation.CompletedAt(); ok {
		_spec.SetField(generation.FieldCompletedAt, field.TypeTime, value)
		_node.CompletedAt = &value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(generation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(generation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := gc.mutation.DeviceInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.DeviceInfoTable,
			Columns: []string{generation.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deviceinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DeviceInfoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.SchedulerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.SchedulerTable,
			Columns: []string{generation.SchedulerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scheduler.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SchedulerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.PromptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.PromptTable,
			Columns: []string{generation.PromptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prompt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PromptID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.NegativePromptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.NegativePromptTable,
			Columns: []string{generation.NegativePromptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: negativeprompt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NegativePromptID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.GenerationModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.GenerationModelTable,
			Columns: []string{generation.GenerationModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationmodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ModelID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generation.UsersTable,
			Columns: []string{generation.UsersColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.GenerationOutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generation.GenerationOutputsTable,
			Columns: []string{generation.GenerationOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GenerationCreateBulk is the builder for creating many Generation entities in bulk.
type GenerationCreateBulk struct {
	config
	builders []*GenerationCreate
}

// Save creates the Generation entities in the database.
func (gcb *GenerationCreateBulk) Save(ctx context.Context) ([]*Generation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Generation, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenerationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenerationCreateBulk) SaveX(ctx context.Context) []*Generation {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GenerationCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GenerationCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
