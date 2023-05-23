// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/user"
)

// ApiToken is the model entity for the ApiToken schema.
type ApiToken struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HashedToken holds the value of the "hashed_token" field.
	HashedToken string `json:"hashed_token,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ShortString holds the value of the "short_string" field.
	ShortString string `json:"short_string,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Uses holds the value of the "uses" field.
	Uses int `json:"uses,omitempty"`
	// CreditsSpent holds the value of the "credits_spent" field.
	CreditsSpent int `json:"credits_spent,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// LastUsedAt holds the value of the "last_used_at" field.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApiTokenQuery when eager-loading is set.
	Edges ApiTokenEdges `json:"edges"`
}

// ApiTokenEdges holds the relations/edges for other nodes in the graph.
type ApiTokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Generations holds the value of the generations edge.
	Generations []*Generation `json:"generations,omitempty"`
	// Upscales holds the value of the upscales edge.
	Upscales []*Upscale `json:"upscales,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApiTokenEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GenerationsOrErr returns the Generations value or an error if the edge
// was not loaded in eager-loading.
func (e ApiTokenEdges) GenerationsOrErr() ([]*Generation, error) {
	if e.loadedTypes[1] {
		return e.Generations, nil
	}
	return nil, &NotLoadedError{edge: "generations"}
}

// UpscalesOrErr returns the Upscales value or an error if the edge
// was not loaded in eager-loading.
func (e ApiTokenEdges) UpscalesOrErr() ([]*Upscale, error) {
	if e.loadedTypes[2] {
		return e.Upscales, nil
	}
	return nil, &NotLoadedError{edge: "upscales"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApiToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apitoken.FieldIsActive:
			values[i] = new(sql.NullBool)
		case apitoken.FieldUses, apitoken.FieldCreditsSpent:
			values[i] = new(sql.NullInt64)
		case apitoken.FieldHashedToken, apitoken.FieldName, apitoken.FieldShortString:
			values[i] = new(sql.NullString)
		case apitoken.FieldLastUsedAt, apitoken.FieldCreatedAt, apitoken.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case apitoken.FieldID, apitoken.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApiToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApiToken fields.
func (at *ApiToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apitoken.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case apitoken.FieldHashedToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_token", values[i])
			} else if value.Valid {
				at.HashedToken = value.String
			}
		case apitoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				at.Name = value.String
			}
		case apitoken.FieldShortString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_string", values[i])
			} else if value.Valid {
				at.ShortString = value.String
			}
		case apitoken.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				at.IsActive = value.Bool
			}
		case apitoken.FieldUses:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uses", values[i])
			} else if value.Valid {
				at.Uses = int(value.Int64)
			}
		case apitoken.FieldCreditsSpent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field credits_spent", values[i])
			} else if value.Valid {
				at.CreditsSpent = int(value.Int64)
			}
		case apitoken.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				at.UserID = *value
			}
		case apitoken.FieldLastUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_at", values[i])
			} else if value.Valid {
				at.LastUsedAt = new(time.Time)
				*at.LastUsedAt = value.Time
			}
		case apitoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = value.Time
			}
		case apitoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				at.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the ApiToken entity.
func (at *ApiToken) QueryUser() *UserQuery {
	return NewApiTokenClient(at.config).QueryUser(at)
}

// QueryGenerations queries the "generations" edge of the ApiToken entity.
func (at *ApiToken) QueryGenerations() *GenerationQuery {
	return NewApiTokenClient(at.config).QueryGenerations(at)
}

// QueryUpscales queries the "upscales" edge of the ApiToken entity.
func (at *ApiToken) QueryUpscales() *UpscaleQuery {
	return NewApiTokenClient(at.config).QueryUpscales(at)
}

// Update returns a builder for updating this ApiToken.
// Note that you need to call ApiToken.Unwrap() before calling this method if this ApiToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *ApiToken) Update() *ApiTokenUpdateOne {
	return NewApiTokenClient(at.config).UpdateOne(at)
}

// Unwrap unwraps the ApiToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *ApiToken) Unwrap() *ApiToken {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApiToken is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *ApiToken) String() string {
	var builder strings.Builder
	builder.WriteString("ApiToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("hashed_token=")
	builder.WriteString(at.HashedToken)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(at.Name)
	builder.WriteString(", ")
	builder.WriteString("short_string=")
	builder.WriteString(at.ShortString)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", at.IsActive))
	builder.WriteString(", ")
	builder.WriteString("uses=")
	builder.WriteString(fmt.Sprintf("%v", at.Uses))
	builder.WriteString(", ")
	builder.WriteString("credits_spent=")
	builder.WriteString(fmt.Sprintf("%v", at.CreditsSpent))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", at.UserID))
	builder.WriteString(", ")
	if v := at.LastUsedAt; v != nil {
		builder.WriteString("last_used_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(at.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(at.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ApiTokens is a parsable slice of ApiToken.
type ApiTokens []*ApiToken

func (at ApiTokens) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}
