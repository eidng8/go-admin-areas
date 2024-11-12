package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	ee "github.com/eidng8/go-ent"
	"github.com/eidng8/go-ent/simpletree"
	"github.com/eidng8/go-ent/softdelete"
	"github.com/ogen-go/ogen"

	gen "github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/intercept"
)

type AdminArea struct {
	ent.Schema
}

func (AdminArea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "admin_areas",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
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
	return append(
		[]ent.Field{
			field.Uint32("id").Unique().Immutable().Annotations(
				// adds constraints to the generated OpenAPI specification
				entoas.Schema(
					&ogen.Schema{
						Type:    "integer",
						Format:  "uint32",
						Minimum: n1,
					},
				),
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
		},
		ee.Timestamps()...,
	)
}

func (AdminArea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Comment out these when running `go generate` for the first time
		softdelete.Mixin{},
		simpletree.ParentU32Mixin[AdminArea]{},
	}
}

func (AdminArea) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		// Comment out this when running `go generate` for the first time
		softdelete.Interceptor(intercept.NewQuery),
	}
}

func (AdminArea) Hooks() []ent.Hook {
	return []ent.Hook{
		// Comment out this when running `go generate` for the first time
		softdelete.Mutator[*gen.Client](),
	}
}
