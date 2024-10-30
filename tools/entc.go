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

func main() {
	oas, err := newOasExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ext := entc.Extensions(oas)
	if err = entc.Generate("./ent/schema", genConfig(), ext); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func newOasExtension() (*entoas.Extension, error) {
	return entoas.NewExtension(
		entoas.Mutations(
			func(g *gen.Graph, s *ogen.Spec) error {
				constraintRequestBody(s.Paths)
				ep := s.Paths["/admin-areas"]
				fixPerPageParamName(ep.Get.Parameters)
				ep.Get.AddParameters(nameParam(), abbrParam(), trashedParam())
				removeEdges(ep.Post)
				setPaginateResponse(ep.Get)
				ep = s.Paths["/admin-areas/{id}"]
				ep.Get.AddParameters(trashedParam())
				removeEdges(ep.Patch)
				ep = s.Paths["/admin-areas/{id}/parent"]
				ep.Get.AddParameters(trashedParam())
				ep = s.Paths["/admin-areas/{id}/children"]
				fixPerPageParamName(ep.Get.Parameters)
				setPaginateResponse(ep.Get)
				ep.Get.AddParameters(nameParam(), abbrParam(), trashedParam())
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
			gen.FeatureVersionedMigration,
		},
	}
}

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

func fixPerPageParamName(params []*ogen.Parameter) {
	for _, param := range params {
		if "itemsPerPage" == param.Name {
			param.Name = "per_page"
		}
	}
}

func removeEdges(op *ogen.Operation) {
	schema := op.RequestBody.Content["application/json"].Schema
	schema.Properties = removeFields(schema.Properties, "parent", "children")
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

func trashedParam() *ogen.Parameter {
	return &ogen.Parameter{
		Name:        "trashed",
		In:          "query",
		Description: "Whether to include trashed items",
		Required:    false,
		Schema:      &ogen.Schema{Type: "boolean"},
	}
}

func setPaginateResponse(op *ogen.Operation) {
	op.Responses["200"] = &ogen.Response{
		Description: "Paginated list of administrative areas",
		Content: map[string]ogen.Media{
			"application/json": {
				Schema: &ogen.Schema{
					Type: "object",
					Properties: []ogen.Property{
						{
							Name: "current_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Page number (1-based)",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "total",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Total number of items",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "per_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Number of items per page",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "last_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Last page number",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "from",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Index (1-based) of the first item in the current page",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "to",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Index (1-based) of the last item in the current page",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "first_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the first page",
							},
						},
						{
							Name: "last_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the last page",
							},
						},
						{
							Name: "next_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the next page",
							},
						},
						{
							Name: "prev_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the previous page",
							},
						},
						{
							Name: "path",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "Base path of the request",
							},
						},
						{
							Name: "data",
							Schema: &ogen.Schema{
								Type:        "array",
								Description: "List of administrative areas",
								Items: &ogen.Items{
									Item: &ogen.Schema{Ref: "#/components/schemas/AdminAreaList"},
								},
							},
						},
					},
					Required: []string{
						"current_page",
						"total",
						"per_page",
						"last_page",
						"from",
						"to",
						"first_page_url",
						"last_page_url",
						"next_page_url",
						"prev_page_url",
						"path",
						"data",
					},
				},
			},
		},
	}
}
