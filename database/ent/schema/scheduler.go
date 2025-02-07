package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Scheduler holds the schema definition for the Scheduler entity.
type Scheduler struct {
	ent.Schema
}

// Fields of the Scheduler.
func (Scheduler) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name_in_worker"),
		field.Bool("is_active").Default(true),
		field.Bool("is_default").Default(false),
		field.Bool("is_hidden").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Scheduler.
func (Scheduler) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M with generations
		edge.To("generations", Generation.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Annotations of the Upscale.
func (Scheduler) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "schedulers"},
	}
}
