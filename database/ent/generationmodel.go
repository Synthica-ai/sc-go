// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generationmodel"
)

// GenerationModel is the model entity for the GenerationModel schema.
type GenerationModel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// NameInWorker holds the value of the "name_in_worker" field.
	NameInWorker string `json:"name_in_worker,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenerationModelQuery when eager-loading is set.
	Edges GenerationModelEdges `json:"edges"`
}

// GenerationModelEdges holds the relations/edges for other nodes in the graph.
type GenerationModelEdges struct {
	// Generations holds the value of the generations edge.
	Generations []*Generation `json:"generations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GenerationsOrErr returns the Generations value or an error if the edge
// was not loaded in eager-loading.
func (e GenerationModelEdges) GenerationsOrErr() ([]*Generation, error) {
	if e.loadedTypes[0] {
		return e.Generations, nil
	}
	return nil, &NotLoadedError{edge: "generations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GenerationModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case generationmodel.FieldNameInWorker:
			values[i] = new(sql.NullString)
		case generationmodel.FieldCreatedAt, generationmodel.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case generationmodel.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GenerationModel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GenerationModel fields.
func (gm *GenerationModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case generationmodel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gm.ID = *value
			}
		case generationmodel.FieldNameInWorker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_in_worker", values[i])
			} else if value.Valid {
				gm.NameInWorker = value.String
			}
		case generationmodel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case generationmodel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gm.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGenerations queries the "generations" edge of the GenerationModel entity.
func (gm *GenerationModel) QueryGenerations() *GenerationQuery {
	return NewGenerationModelClient(gm.config).QueryGenerations(gm)
}

// Update returns a builder for updating this GenerationModel.
// Note that you need to call GenerationModel.Unwrap() before calling this method if this GenerationModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GenerationModel) Update() *GenerationModelUpdateOne {
	return NewGenerationModelClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GenerationModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GenerationModel) Unwrap() *GenerationModel {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GenerationModel is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GenerationModel) String() string {
	var builder strings.Builder
	builder.WriteString("GenerationModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("name_in_worker=")
	builder.WriteString(gm.NameInWorker)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gm.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GenerationModels is a parsable slice of GenerationModel.
type GenerationModels []*GenerationModel

func (gm GenerationModels) config(cfg config) {
	for _i := range gm {
		gm[_i].config = cfg
	}
}
