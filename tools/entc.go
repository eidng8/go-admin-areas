//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func removeFields(props []ogen.Property, fields ...string) []ogen.Property {
	for _, field := range fields {
		for i, prop := range props {
			if prop.Name == field {
				props = append(props[:i], props[i+1:]...)
				break
			}
		}
	}
	return props

}

func main() {
	ex, err := entoas.NewExtension(
		entoas.Mutations(
			func(g *gen.Graph, s *ogen.Spec) error {
				// Is there any way to do this in schema?
				schema := s.Paths["/admin-areas"].Post.RequestBody.Content["application/json"].Schema
				schema.Properties = removeFields(
					schema.Properties, "parent", "children",
				)
				schema = s.Paths["/admin-areas/{id}"].Patch.RequestBody.Content["application/json"].Schema
				schema.Properties = removeFields(
					schema.Properties, "parent", "children",
				)
				// Also expose PUT as PATCH!?
				// Don't do this as it will generate 2 UpdateAdminArea functions
				// that will cause duplicate function definition error
				//s.Paths["/admin-areas/{id}"].Put = s.Paths["/admin-areas/{id}"].Patch
				return nil
			},
		),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate(
		"./ent/schema", &gen.Config{
			Features: []gen.Feature{
				gen.FeatureIntercept,
				gen.FeatureSnapshot,
			},
		}, entc.Extensions(ex),
	)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
