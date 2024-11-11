//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-ent/simpletree"
	"github.com/eidng8/go-ent/softdelete"
	"github.com/ogen-go/ogen"
)

func main() {
	oas, err := newOasExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ext := entc.Extensions(oas, &simpletree.Extension{})
	err = entc.Generate("./ent/schema", genConfig(), ext)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func newOasExtension() (*entoas.Extension, error) {
	return entoas.NewExtension(
		entoas.Mutations(
			func(g *gen.Graph, s *ogen.Spec) error {
				genSpec(s)
				constraintRequestBody(s.Paths)
				softdelete.AttachTo(s, "/admin-areas", "AdminAreaRead")
				ep := s.Paths["/admin-areas"]
				ep.Get.AddParameters(nameParam(), abbrParam())
				simpletree.RemoveEdges(ep.Post)
				paginate.FixParamNames(ep.Get.Parameters)
				paginate.SetResponse(
					ep.Get, "Paginated list of administrative areas",
					"#/components/schemas/AdminAreaList",
				)
				ep = s.Paths["/admin-areas/{id}"]
				simpletree.RemoveEdges(ep.Patch)
				ep = s.Paths["/admin-areas/{id}/children"]
				ep.Get.AddParameters(nameParam(), abbrParam())
				paginate.FixParamNames(ep.Get.Parameters)
				paginate.SetResponse(
					ep.Get,
					"Paginated list of subordinate administrative areas. Pagination is disabled when `recurse` is true.",
					"#/components/schemas/AdminAreaList",
				)
				simpletree.AttachTo(ep.Get)
				return nil
			},
		),
	)
}

func genConfig() *gen.Config {
	return &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureExecQuery,
			gen.FeatureVersionedMigration,
		},
	}
}

func genSpec(s *ogen.Spec) {
	s.Info.SetTitle("Administrative areas listing API")
	s.Info.SetDescription("This is an API listing administrative areas")
	s.Info.SetVersion("0.0.1")
}

func constraintRequestBody(paths ogen.Paths) {
	b := false
	for _, path := range paths {
		for _, op := range []*ogen.Operation{path.Put, path.Post, path.Patch} {
			if nil == op || nil == op.RequestBody || nil == op.RequestBody.Content {
				continue
			}
			for _, param := range op.RequestBody.Content {
				if nil == param.Schema {
					continue
				}
				param.Schema.AdditionalProperties = &ogen.AdditionalProperties{Bool: &b}
			}
		}
	}
}

func nameParam() *ogen.Parameter {
	u2 := uint64(2)
	u255 := uint64(255)
	return &ogen.Parameter{
		Name:        "name",
		In:          "query",
		Description: "Name of the administrative area",
		Required:    false,
		Schema: &ogen.Schema{
			Type:      "string",
			MinLength: &u2,
			MaxLength: &u255,
		},
	}
}

func abbrParam() *ogen.Parameter {
	u2 := uint64(2)
	u255 := uint64(255)
	return &ogen.Parameter{
		Name:        "abbr",
		In:          "query",
		Description: "Abbreviation of the administrative area, can be a CSV list",
		Required:    false,
		Schema: &ogen.Schema{
			Type:      "string",
			MinLength: &u2,
			MaxLength: &u255,
		},
	}
}
