package schema

import (
	"time"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"
)

type AdminArea struct {
	ent.Schema
}

func (AdminArea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "admin_areas"},
		entsql.WithComments(true),
		entsql.OnDelete(entsql.Restrict),
		schema.Comment("Administrative area table"),
	}
}

func (AdminArea) Fields() []ent.Field {
	n1 := ogen.Num("1")
	u1 := uint64(1)
	u2 := uint64(2)
	u255 := uint64(255)
	return []ent.Field{
		field.Uint32("id").Unique().Immutable().Annotations(
			// adds constraints to the generated OpenAPI specification
			entoas.Schema(&ogen.Schema{Type: "integer", Minimum: n1}),
		),
		field.String("name").NotEmpty().MinLen(2).MaxLen(255).
			Comment("Administrative area name").Annotations(
			// adds constraints to the generated OpenAPI specification
			entoas.Schema(
				&ogen.Schema{
					Type:        "string",
					MinLength:   &u2,
					MaxLength:   &u255,
					Description: "Administrative area name",
				},
			),
		),
		field.String("abbr").Optional().Nillable().MinLen(1).MaxLen(255).
			Comment("Administrative area abbreviation, CSV values").Annotations(
			entoas.Schema(
				// adds constraints to the generated OpenAPI specification
				&ogen.Schema{
					Type:        "string",
					MinLength:   &u1,
					MaxLength:   &u255,
					Nullable:    true,
					Description: "Administrative area abbreviations, CSV values",
				},
			),
		),
		field.Text("memo").Optional().Nillable().Comment("Remarks").
			Comment("Remarks").Annotations(
			entoas.Schema(
				// adds constraints to the generated OpenAPI specification
				&ogen.Schema{
					Type:        "string",
					MinLength:   &u1,
					MaxLength:   &u255,
					Nullable:    true,
					Description: "Remarks",
				},
			),
		),
		field.Time("created_at").Optional().Nillable().Default(time.Now).
			Immutable().Annotations(
			// removes the field from the create & update OpenAPI endpoint
			entoas.ReadOnly(true),
		),
		field.Time("updated_at").Optional().Nillable().UpdateDefault(time.Now).
			Immutable().Annotations(
			// removes the field from the create & update OpenAPI endpoint
			entoas.ReadOnly(true),
		),
		field.Int("lft").Optional().Nillable().Annotations(
			// skips the field from OpenAPI specification
			entoas.Skip(true),
			// adds constraints to the generated OpenAPI specification
			entoas.Schema(&ogen.Schema{Type: "integer", Minimum: n1}),
		),
		field.Int("rgt").Optional().Nillable().Annotations(
			// skips the field from OpenAPI specification
			entoas.Skip(true),
			// adds constraints to the generated OpenAPI specification
			entoas.Schema(&ogen.Schema{Type: "integer", Minimum: n1}),
		),
		field.Uint32("parent_id").Optional().Nillable().Annotations(
			entoas.Schema(
				&ogen.Schema{
					Type:    "integer",
					Minimum: n1,
				},
			),
		),
	}
}

func (AdminArea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", AdminArea.Type).
			Annotations(
				entsql.OnDelete(entsql.Restrict),
				entoas.ReadOnly(true),
				entoas.Skip(true),
			).
			From("parent").Field("parent_id").Unique(),
	}
}

func (AdminArea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		//PaginationMixin{},
	}
}
