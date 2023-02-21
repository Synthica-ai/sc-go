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
	"github.com/stablecog/sc-go/database/ent/credit"
	"github.com/stablecog/sc-go/database/ent/credittype"
	"github.com/stablecog/sc-go/database/ent/user"
)

// CreditCreate is the builder for creating a Credit entity.
type CreditCreate struct {
	config
	mutation *CreditMutation
	hooks    []Hook
}

// SetRemainingAmount sets the "remaining_amount" field.
func (cc *CreditCreate) SetRemainingAmount(i int32) *CreditCreate {
	cc.mutation.SetRemainingAmount(i)
	return cc
}

// SetExpiresAt sets the "expires_at" field.
func (cc *CreditCreate) SetExpiresAt(t time.Time) *CreditCreate {
	cc.mutation.SetExpiresAt(t)
	return cc
}

// SetStripeLineItemID sets the "stripe_line_item_id" field.
func (cc *CreditCreate) SetStripeLineItemID(s string) *CreditCreate {
	cc.mutation.SetStripeLineItemID(s)
	return cc
}

// SetNillableStripeLineItemID sets the "stripe_line_item_id" field if the given value is not nil.
func (cc *CreditCreate) SetNillableStripeLineItemID(s *string) *CreditCreate {
	if s != nil {
		cc.SetStripeLineItemID(*s)
	}
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CreditCreate) SetUserID(u uuid.UUID) *CreditCreate {
	cc.mutation.SetUserID(u)
	return cc
}

// SetCreditTypeID sets the "credit_type_id" field.
func (cc *CreditCreate) SetCreditTypeID(u uuid.UUID) *CreditCreate {
	cc.mutation.SetCreditTypeID(u)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CreditCreate) SetCreatedAt(t time.Time) *CreditCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CreditCreate) SetNillableCreatedAt(t *time.Time) *CreditCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CreditCreate) SetUpdatedAt(t time.Time) *CreditCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CreditCreate) SetNillableUpdatedAt(t *time.Time) *CreditCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CreditCreate) SetID(u uuid.UUID) *CreditCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CreditCreate) SetNillableID(u *uuid.UUID) *CreditCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cc *CreditCreate) SetUsersID(id uuid.UUID) *CreditCreate {
	cc.mutation.SetUsersID(id)
	return cc
}

// SetUsers sets the "users" edge to the User entity.
func (cc *CreditCreate) SetUsers(u *User) *CreditCreate {
	return cc.SetUsersID(u.ID)
}

// SetCreditType sets the "credit_type" edge to the CreditType entity.
func (cc *CreditCreate) SetCreditType(c *CreditType) *CreditCreate {
	return cc.SetCreditTypeID(c.ID)
}

// Mutation returns the CreditMutation object of the builder.
func (cc *CreditCreate) Mutation() *CreditMutation {
	return cc.mutation
}

// Save creates the Credit in the database.
func (cc *CreditCreate) Save(ctx context.Context) (*Credit, error) {
	cc.defaults()
	return withHooks[*Credit, CreditMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CreditCreate) SaveX(ctx context.Context) *Credit {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CreditCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CreditCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CreditCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := credit.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := credit.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := credit.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CreditCreate) check() error {
	if _, ok := cc.mutation.RemainingAmount(); !ok {
		return &ValidationError{Name: "remaining_amount", err: errors.New(`ent: missing required field "Credit.remaining_amount"`)}
	}
	if _, ok := cc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Credit.expires_at"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Credit.user_id"`)}
	}
	if _, ok := cc.mutation.CreditTypeID(); !ok {
		return &ValidationError{Name: "credit_type_id", err: errors.New(`ent: missing required field "Credit.credit_type_id"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Credit.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Credit.updated_at"`)}
	}
	if _, ok := cc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Credit.users"`)}
	}
	if _, ok := cc.mutation.CreditTypeID(); !ok {
		return &ValidationError{Name: "credit_type", err: errors.New(`ent: missing required edge "Credit.credit_type"`)}
	}
	return nil
}

func (cc *CreditCreate) sqlSave(ctx context.Context) (*Credit, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CreditCreate) createSpec() (*Credit, *sqlgraph.CreateSpec) {
	var (
		_node = &Credit{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: credit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: credit.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.RemainingAmount(); ok {
		_spec.SetField(credit.FieldRemainingAmount, field.TypeInt32, value)
		_node.RemainingAmount = value
	}
	if value, ok := cc.mutation.ExpiresAt(); ok {
		_spec.SetField(credit.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := cc.mutation.StripeLineItemID(); ok {
		_spec.SetField(credit.FieldStripeLineItemID, field.TypeString, value)
		_node.StripeLineItemID = &value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(credit.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(credit.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credit.UsersTable,
			Columns: []string{credit.UsersColumn},
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
	if nodes := cc.mutation.CreditTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credit.CreditTypeTable,
			Columns: []string{credit.CreditTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credittype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreditTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CreditCreateBulk is the builder for creating many Credit entities in bulk.
type CreditCreateBulk struct {
	config
	builders []*CreditCreate
}

// Save creates the Credit entities in the database.
func (ccb *CreditCreateBulk) Save(ctx context.Context) ([]*Credit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Credit, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CreditMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CreditCreateBulk) SaveX(ctx context.Context) []*Credit {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CreditCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CreditCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
