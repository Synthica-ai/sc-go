// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/credit"
	"github.com/stablecog/sc-go/database/ent/credittype"
	"github.com/stablecog/sc-go/database/ent/user"
)

// Credit is the model entity for the Credit schema.
type Credit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// RemainingAmount holds the value of the "remaining_amount" field.
	RemainingAmount int32 `json:"remaining_amount,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// StripeLineItemID holds the value of the "stripe_line_item_id" field.
	StripeLineItemID *string `json:"stripe_line_item_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CreditTypeID holds the value of the "credit_type_id" field.
	CreditTypeID uuid.UUID `json:"credit_type_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CreditQuery when eager-loading is set.
	Edges CreditEdges `json:"edges"`
}

// CreditEdges holds the relations/edges for other nodes in the graph.
type CreditEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// CreditType holds the value of the credit_type edge.
	CreditType *CreditType `json:"credit_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CreditEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// CreditTypeOrErr returns the CreditType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CreditEdges) CreditTypeOrErr() (*CreditType, error) {
	if e.loadedTypes[1] {
		if e.CreditType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: credittype.Label}
		}
		return e.CreditType, nil
	}
	return nil, &NotLoadedError{edge: "credit_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Credit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case credit.FieldRemainingAmount:
			values[i] = new(sql.NullInt64)
		case credit.FieldStripeLineItemID:
			values[i] = new(sql.NullString)
		case credit.FieldExpiresAt, credit.FieldCreatedAt, credit.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case credit.FieldID, credit.FieldUserID, credit.FieldCreditTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Credit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Credit fields.
func (c *Credit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case credit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case credit.FieldRemainingAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remaining_amount", values[i])
			} else if value.Valid {
				c.RemainingAmount = int32(value.Int64)
			}
		case credit.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				c.ExpiresAt = value.Time
			}
		case credit.FieldStripeLineItemID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stripe_line_item_id", values[i])
			} else if value.Valid {
				c.StripeLineItemID = new(string)
				*c.StripeLineItemID = value.String
			}
		case credit.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				c.UserID = *value
			}
		case credit.FieldCreditTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field credit_type_id", values[i])
			} else if value != nil {
				c.CreditTypeID = *value
			}
		case credit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case credit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Credit entity.
func (c *Credit) QueryUsers() *UserQuery {
	return NewCreditClient(c.config).QueryUsers(c)
}

// QueryCreditType queries the "credit_type" edge of the Credit entity.
func (c *Credit) QueryCreditType() *CreditTypeQuery {
	return NewCreditClient(c.config).QueryCreditType(c)
}

// Update returns a builder for updating this Credit.
// Note that you need to call Credit.Unwrap() before calling this method if this Credit
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Credit) Update() *CreditUpdateOne {
	return NewCreditClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Credit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Credit) Unwrap() *Credit {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Credit is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Credit) String() string {
	var builder strings.Builder
	builder.WriteString("Credit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("remaining_amount=")
	builder.WriteString(fmt.Sprintf("%v", c.RemainingAmount))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(c.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.StripeLineItemID; v != nil {
		builder.WriteString("stripe_line_item_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", ")
	builder.WriteString("credit_type_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CreditTypeID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Credits is a parsable slice of Credit.
type Credits []*Credit

func (c Credits) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
